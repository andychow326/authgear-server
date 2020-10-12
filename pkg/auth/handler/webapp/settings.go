package webapp

import (
	"net/http"

	"github.com/authgear/authgear-server/pkg/auth/handler/webapp/viewmodels"
	"github.com/authgear/authgear-server/pkg/lib/authn"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
	"github.com/authgear/authgear-server/pkg/lib/session"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/httputil"
	"github.com/authgear/authgear-server/pkg/util/template"
)

const (
	TemplateItemTypeAuthUISettingsHTML string = "auth_ui_settings.html"
)

var TemplateAuthUISettingsHTML = template.Register(template.T{
	Type:                    TemplateItemTypeAuthUISettingsHTML,
	IsHTML:                  true,
	TranslationTemplateType: TemplateItemTypeAuthUITranslationJSON,
	Defines:                 defines,
	ComponentTemplateTypes:  components,
})

func ConfigureSettingsRoute(route httproute.Route) httproute.Route {
	return route.
		WithMethods("OPTIONS", "GET", "POST").
		WithPathPattern("/settings")
}

type SettingsViewModel struct {
	Authenticators           []*authenticator.Info
	SecondaryTOTPAllowed     bool
	SecondaryOOBOTPAllowed   bool
	SecondaryPasswordAllowed bool
}

type SettingsAuthenticatorService interface {
	List(userID string, filters ...authenticator.Filter) ([]*authenticator.Info, error)
}

type SettingsMFAService interface {
	InvalidateAllDeviceTokens(userID string) error
}

type SettingsHandler struct {
	Database       *db.Handle
	BaseViewModel  *viewmodels.BaseViewModeler
	Renderer       Renderer
	Authentication *config.AuthenticationConfig
	Authenticators SettingsAuthenticatorService
	MFA            SettingsMFAService
}

func (h *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	redirectURI := httputil.HostRelative(r.URL).String()
	userID := *session.GetUserID(r.Context())

	if r.Method == "GET" {
		data := map[string]interface{}{}

		baseViewModel := h.BaseViewModel.ViewModel(r, nil)
		viewmodels.Embed(data, baseViewModel)

		authenticators, err := h.Authenticators.List(userID)
		if err != nil {
			panic(err)
		}

		totp := false
		oobotp := false
		password := false
		for _, typ := range h.Authentication.SecondaryAuthenticators {
			switch typ {
			case authn.AuthenticatorTypePassword:
				password = true
			case authn.AuthenticatorTypeTOTP:
				totp = true
			case authn.AuthenticatorTypeOOB:
				oobotp = true
			}
		}

		viewModel := SettingsViewModel{
			Authenticators:           authenticators,
			SecondaryTOTPAllowed:     totp,
			SecondaryOOBOTPAllowed:   oobotp,
			SecondaryPasswordAllowed: password,
		}
		viewmodels.Embed(data, viewModel)

		h.Renderer.RenderHTML(w, r, TemplateItemTypeAuthUISettingsHTML, data)
		return
	}

	if r.Method == "POST" && r.Form.Get("x_action") == "revoke_devices" {
		err := h.Database.WithTx(func() error {
			if err := h.MFA.InvalidateAllDeviceTokens(userID); err != nil {
				return err
			}
			http.Redirect(w, r, redirectURI, http.StatusFound)
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
}
