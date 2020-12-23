package webapp

import (
	"context"
	"time"

	"github.com/authgear/authgear-server/pkg/util/base32"
	corerand "github.com/authgear/authgear-server/pkg/util/rand"
)

type sessionContextKey struct{}

func GetSession(ctx context.Context) *Session {
	s, _ := ctx.Value(sessionContextKey{}).(*Session)
	return s
}

func WithSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionContextKey{}, session)
}

type SessionOptions struct {
	WsChannelID     string
	RedirectURI     string
	KeepAfterFinish bool
	UILocales       string
	Prompt          string
	SDK             string
	Extra           map[string]interface{}
	UpdatedAt       time.Time
}

func NewSessionOptionsFromSession(s *Session) SessionOptions {
	return SessionOptions{
		WsChannelID:     s.WsChannelID,
		RedirectURI:     s.RedirectURI,
		KeepAfterFinish: s.KeepAfterFinish,
		UILocales:       s.UILocales,
		Prompt:          s.Prompt,
		SDK:             s.SDK,
		Extra:           nil, // Omit extra by default
	}
}

type Session struct {
	ID string `json:"id"`

	WsChannelID string `json:"ws_channel_id"`

	// Steps is a history stack of steps taken within this session.
	Steps []SessionStep `json:"steps,omitempty"`

	// RedirectURI is the URI to redirect to after the completion of session.
	RedirectURI string `json:"redirect_uri,omitempty"`

	// KeepAfterFinish indicates the session would not be deleted after the
	// completion of interaction graph.
	KeepAfterFinish bool `json:"keep_after_finish,omitempty"`

	// Extra is used to store extra information for use of webapp.
	Extra map[string]interface{} `json:"extra"`

	// Prompt is used to indicate requested authentication behavior
	Prompt string `json:"prompt,omitempty"`

	// SDK is used to indicate the request is triggered by sdk
	SDK string `json:"sdk,omitempty"`

	// UILocales are the locale to be used to render UI, passed in from OAuth
	// flow or query parameter.
	UILocales string `json:"ui_locales,omitempty"`

	// UpdatedAt indicate the session last updated time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func newSessionID() string {
	const (
		idAlphabet string = base32.Alphabet
		idLength   int    = 32
	)
	return corerand.StringWithAlphabet(idLength, idAlphabet, corerand.SecureRand)
}

func NewSession(options SessionOptions) *Session {
	id := newSessionID()
	if options.WsChannelID == "" {
		options.WsChannelID = id
	}
	s := &Session{
		ID:              id,
		WsChannelID:     options.WsChannelID,
		RedirectURI:     options.RedirectURI,
		KeepAfterFinish: options.KeepAfterFinish,
		Extra:           make(map[string]interface{}),
		Prompt:          options.Prompt,
		SDK:             options.SDK,
		UILocales:       options.UILocales,
		UpdatedAt:       options.UpdatedAt,
	}
	for k, v := range options.Extra {
		s.Extra[k] = v
	}
	if s.RedirectURI == "" {
		s.RedirectURI = "/"
	}
	return s
}

func (s *Session) CurrentStep() SessionStep {
	return s.Steps[len(s.Steps)-1]
}
