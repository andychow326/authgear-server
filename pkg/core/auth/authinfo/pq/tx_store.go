package pq

import (
	"github.com/sirupsen/logrus"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	"github.com/skygeario/skygear-server/pkg/core/db"
)

type safeAuthInfoStore struct {
	impl      *authInfoStore
	txContext db.SafeTxContext
}

func NewSafeAuthInfoStore(
	builder db.SQLBuilder,
	executor db.SQLExecutor,
	logger *logrus.Entry,
	txContext db.SafeTxContext,
) authinfo.Store {
	return &safeAuthInfoStore{
		impl:      newAuthInfoStore(builder, executor, logger),
		txContext: txContext,
	}
}

func (s *safeAuthInfoStore) CreateAuth(authinfo *authinfo.AuthInfo) error {
	s.txContext.EnsureTx()
	return s.impl.CreateAuth(authinfo)
}

func (s *safeAuthInfoStore) GetAuth(id string, authinfo *authinfo.AuthInfo) error {
	s.txContext.EnsureTx()
	return s.impl.GetAuth(id, authinfo)
}

func (s *safeAuthInfoStore) UpdateAuth(authinfo *authinfo.AuthInfo) error {
	s.txContext.EnsureTx()
	return s.impl.UpdateAuth(authinfo)
}

func (s *safeAuthInfoStore) DeleteAuth(id string) error {
	s.txContext.EnsureTx()
	return s.impl.DeleteAuth(id)
}

// this ensures that our structure conform to certain interfaces.
var (
	_ authinfo.Store = &safeAuthInfoStore{}
)
