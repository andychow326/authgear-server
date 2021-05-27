package tenant

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/util/errorutil"
	"github.com/authgear/authgear-server/pkg/util/log"
)

type Handle struct {
	ctx         context.Context
	pool        *Pool
	cfg         *config.DatabaseConfig
	credentials *config.DatabaseCredentials
	logger      *log.Logger

	tx    *sqlx.Tx
	hooks []TransactionHook
}

func NewHandle(ctx context.Context, pool *Pool, cfg *config.DatabaseConfig, credentials *config.DatabaseCredentials, lf *log.Factory) *Handle {
	return &Handle{
		ctx:         ctx,
		pool:        pool,
		cfg:         cfg,
		credentials: credentials,
		logger:      lf.New("db-handle"),
	}
}

func (h *Handle) Conn() (sqlx.ExtContext, error) {
	tx := h.tx
	if tx == nil {
		panic("db: transaction not started")
	}
	return tx, nil
}

func (h *Handle) UseHook(hook TransactionHook) {
	h.hooks = append(h.hooks, hook)
}

// WithTx commits if do finishes without error and rolls back otherwise.
func (h *Handle) WithTx(do func() error) (err error) {
	if err = h.beginTx(); err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			_ = h.rollbackTx()
			panic(r)
		} else if err != nil {
			if rbErr := h.rollbackTx(); rbErr != nil {
				h.logger.WithError(rbErr).Error("failed to rollback tx")
			}
		} else {
			err = h.commitTx()
		}
	}()

	err = do()
	return
}

// ReadOnly runs do in a transaction and rolls back always.
func (h *Handle) ReadOnly(do func() error) (err error) {
	if err = h.beginTx(); err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			_ = h.rollbackTx()
			panic(r)
		} else if err != nil {
			if rbErr := h.rollbackTx(); rbErr != nil {
				h.logger.WithError(rbErr).Error("failed to rollback tx")
			}
		} else {
			err = h.rollbackTx()
		}
	}()

	err = do()
	return
}

func (h *Handle) beginTx() error {
	if h.tx != nil {
		panic("db: a transaction has already begun")
	}

	db, err := h.openDB()
	if err != nil {
		return err
	}
	tx, err := db.BeginTxx(h.ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	h.tx = tx

	return nil
}

func (h *Handle) commitTx() error {
	if h.tx == nil {
		panic("db: a transaction has not begun")
	}

	for _, hook := range h.hooks {
		err := hook.WillCommitTx()
		if err != nil {
			if rbErr := h.tx.Rollback(); rbErr != nil {
				err = errorutil.WithSecondaryError(err, rbErr)
			}
			return err
		}
	}

	err := h.tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	h.tx = nil

	for _, hook := range h.hooks {
		hook.DidCommitTx()
	}

	return nil
}

func (h *Handle) rollbackTx() error {
	if h.tx == nil {
		panic("db: a transaction has not begun")
	}

	err := h.tx.Rollback()
	if err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}

	h.tx = nil
	return nil
}

func (h *Handle) openDB() (*sqlx.DB, error) {
	opts := OpenOptions{
		URL:             h.credentials.DatabaseURL,
		MaxOpenConns:    *h.cfg.MaxOpenConnection,
		MaxIdleConns:    *h.cfg.MaxIdleConnection,
		ConnMaxLifetime: h.cfg.MaxConnectionLifetime.Duration(),
		ConnMaxIdleTime: h.cfg.IdleConnectionTimeout.Duration(),
	}
	h.logger.WithFields(map[string]interface{}{
		"max_open_conns":             opts.MaxOpenConns,
		"max_idle_conns":             opts.MaxIdleConns,
		"conn_max_lifetime_seconds":  opts.ConnMaxLifetime.Seconds(),
		"conn_max_idle_time_seconds": opts.ConnMaxIdleTime.Seconds(),
	}).Debug("open database")

	db, err := h.pool.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
