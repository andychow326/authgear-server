package sms

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/hook"
)

var ErrMissingCustomSMSProviderConfiguration = errors.New("sms: custom provider configuration is missing")

type CustomClient struct {
	Config   *config.CustomSMSProviderConfigs
	DenoHook hook.DenoHook
}

func NewCustomClient(c *config.CustomSMSProviderConfigs, d hook.DenoHook) *CustomClient {
	if c == nil {
		return nil
	}

	return &CustomClient{
		Config:   c,
		DenoHook: d,
	}
}

type SendSMSPayload struct {
	To   string `json:"to"`
	Body string `json:"body"`
}

func (c *CustomClient) Send(from string, to string, body string) error {
	if c.Config == nil {
		return ErrMissingCustomSMSProviderConfiguration
	}
	u, err := url.Parse(c.Config.URL)
	if err != nil {
		return err
	}
	switch {
	case c.DenoHook.SupportURL(u):
		_, err := c.DenoHook.RunSync(u, &SendSMSPayload{To: to, Body: body})
		return err
	default:
		return fmt.Errorf("unsupported hook URL: %v", u)
	}
}
