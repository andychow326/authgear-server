package deps

import (
	"net/http"

	"github.com/google/wire"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/interaction"
	"github.com/authgear/authgear-server/pkg/lib/session/idpsession"
	"github.com/authgear/authgear-server/pkg/util/httputil"
)

var utilsDeps = wire.NewSet(
	wire.NewSet(
		NewCookieFactory,
		wire.Bind(new(idpsession.CookieFactory), new(*httputil.CookieFactory)),
		wire.Bind(new(interaction.CookieFactory), new(*httputil.CookieFactory)),
	),
)

func NewCookieFactory(r *http.Request, serverConfig *config.ServerConfig) *httputil.CookieFactory {
	return &httputil.CookieFactory{
		Request:    r,
		TrustProxy: serverConfig.TrustProxy,
	}
}
