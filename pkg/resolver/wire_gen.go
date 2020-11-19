// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package resolver

import (
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/oob"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/password"
	service2 "github.com/authgear/authgear-server/pkg/lib/authn/authenticator/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator/totp"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/anonymous"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/loginid"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/oauth"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity/service"
	"github.com/authgear/authgear-server/pkg/lib/authn/mfa"
	"github.com/authgear/authgear-server/pkg/lib/authn/user"
	"github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/facade"
	"github.com/authgear/authgear-server/pkg/lib/feature/verification"
	"github.com/authgear/authgear-server/pkg/lib/feature/welcomemessage"
	"github.com/authgear/authgear-server/pkg/lib/infra/db"
	"github.com/authgear/authgear-server/pkg/lib/infra/middleware"
	oauth2 "github.com/authgear/authgear-server/pkg/lib/oauth"
	"github.com/authgear/authgear-server/pkg/lib/oauth/oidc"
	"github.com/authgear/authgear-server/pkg/lib/oauth/pq"
	"github.com/authgear/authgear-server/pkg/lib/oauth/redis"
	"github.com/authgear/authgear-server/pkg/lib/ratelimit"
	"github.com/authgear/authgear-server/pkg/lib/session"
	"github.com/authgear/authgear-server/pkg/lib/session/access"
	"github.com/authgear/authgear-server/pkg/lib/session/idpsession"
	"github.com/authgear/authgear-server/pkg/lib/translation"
	"github.com/authgear/authgear-server/pkg/lib/web"
	"github.com/authgear/authgear-server/pkg/resolver/handler"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/rand"
	"github.com/authgear/authgear-server/pkg/util/template"
	"net/http"
)

// Injectors from wire.go:

func newSentryMiddleware(p *deps.RootProvider) httproute.Middleware {
	hub := p.SentryHub
	environmentConfig := p.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	sentryMiddleware := &middleware.SentryMiddleware{
		SentryHub:  hub,
		TrustProxy: trustProxy,
	}
	return sentryMiddleware
}

func newPanicEndMiddleware(p *deps.RootProvider) httproute.Middleware {
	panicEndMiddleware := &middleware.PanicEndMiddleware{}
	return panicEndMiddleware
}

func newPanicWriteEmptyResponseMiddleware(p *deps.RootProvider) httproute.Middleware {
	panicWriteEmptyResponseMiddleware := &middleware.PanicWriteEmptyResponseMiddleware{}
	return panicWriteEmptyResponseMiddleware
}

func newBodyLimitMiddleware(p *deps.RootProvider) httproute.Middleware {
	bodyLimitMiddleware := &middleware.BodyLimitMiddleware{}
	return bodyLimitMiddleware
}

func newPanicLogMiddleware(p *deps.RequestProvider) httproute.Middleware {
	appProvider := p.AppProvider
	factory := appProvider.LoggerFactory
	panicLogMiddlewareLogger := middleware.NewPanicLogMiddlewareLogger(factory)
	panicLogMiddleware := &middleware.PanicLogMiddleware{
		Logger: panicLogMiddlewareLogger,
	}
	return panicLogMiddleware
}

func newSessionMiddleware(p *deps.RequestProvider) httproute.Middleware {
	request := p.Request
	appProvider := p.AppProvider
	rootProvider := appProvider.RootProvider
	environmentConfig := rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	cookieFactory := deps.NewCookieFactory(request, trustProxy)
	config := appProvider.Config
	appConfig := config.AppConfig
	httpConfig := appConfig.HTTP
	sessionConfig := appConfig.Session
	cookieDef := idpsession.NewSessionCookieDef(httpConfig, sessionConfig)
	handle := appProvider.Redis
	appID := appConfig.ID
	clock := _wireSystemClockValue
	factory := appProvider.LoggerFactory
	storeRedisLogger := idpsession.NewStoreRedisLogger(factory)
	storeRedis := &idpsession.StoreRedis{
		Redis:  handle,
		AppID:  appID,
		Clock:  clock,
		Logger: storeRedisLogger,
	}
	eventStoreRedis := &access.EventStoreRedis{
		Redis: handle,
		AppID: appID,
	}
	eventProvider := &access.EventProvider{
		Store: eventStoreRedis,
	}
	rand := _wireRandValue
	provider := &idpsession.Provider{
		Request:      request,
		Store:        storeRedis,
		AccessEvents: eventProvider,
		TrustProxy:   trustProxy,
		Config:       sessionConfig,
		Clock:        clock,
		Random:       rand,
	}
	resolver := &idpsession.Resolver{
		CookieFactory: cookieFactory,
		Cookie:        cookieDef,
		Provider:      provider,
		TrustProxy:    trustProxy,
		Clock:         clock,
	}
	secretConfig := config.SecretConfig
	databaseCredentials := deps.ProvideDatabaseCredentials(secretConfig)
	sqlBuilder := db.ProvideSQLBuilder(databaseCredentials, appID)
	context := deps.ProvideRequestContext(request)
	dbHandle := appProvider.Database
	sqlExecutor := db.SQLExecutor{
		Context:  context,
		Database: dbHandle,
	}
	authorizationStore := &pq.AuthorizationStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	logger := redis.NewLogger(factory)
	grantStore := &redis.GrantStore{
		Redis:       handle,
		AppID:       appID,
		Logger:      logger,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
		Clock:       clock,
	}
	oAuthKeyMaterials := deps.ProvideOAuthKeyMaterials(secretConfig)
	endpointsProvider := &EndpointsProvider{
		HTTP: httpConfig,
	}
	store := &user.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	authenticationConfig := appConfig.Authentication
	identityConfig := appConfig.Identity
	serviceStore := &service.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	loginidStore := &loginid.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	loginIDConfig := identityConfig.LoginID
	manager := appProvider.Resources
	typeCheckerFactory := &loginid.TypeCheckerFactory{
		Config:    loginIDConfig,
		Resources: manager,
	}
	checker := &loginid.Checker{
		Config:             loginIDConfig,
		TypeCheckerFactory: typeCheckerFactory,
	}
	normalizerFactory := &loginid.NormalizerFactory{
		Config: loginIDConfig,
	}
	loginidProvider := &loginid.Provider{
		Store:             loginidStore,
		Config:            loginIDConfig,
		Checker:           checker,
		NormalizerFactory: normalizerFactory,
		Clock:             clock,
	}
	oauthStore := &oauth.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	oauthProvider := &oauth.Provider{
		Store: oauthStore,
		Clock: clock,
	}
	anonymousStore := &anonymous.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	anonymousProvider := &anonymous.Provider{
		Store: anonymousStore,
		Clock: clock,
	}
	serviceService := &service.Service{
		Authentication: authenticationConfig,
		Identity:       identityConfig,
		Store:          serviceStore,
		LoginID:        loginidProvider,
		OAuth:          oauthProvider,
		Anonymous:      anonymousProvider,
	}
	store2 := &service2.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	passwordStore := &password.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	authenticatorConfig := appConfig.Authenticator
	authenticatorPasswordConfig := authenticatorConfig.Password
	passwordLogger := password.NewLogger(factory)
	historyStore := &password.HistoryStore{
		Clock:       clock,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	passwordChecker := password.ProvideChecker(authenticatorPasswordConfig, historyStore)
	housekeeperLogger := password.NewHousekeeperLogger(factory)
	housekeeper := &password.Housekeeper{
		Store:  historyStore,
		Logger: housekeeperLogger,
		Config: authenticatorPasswordConfig,
	}
	passwordProvider := &password.Provider{
		Store:           passwordStore,
		Config:          authenticatorPasswordConfig,
		Clock:           clock,
		Logger:          passwordLogger,
		PasswordHistory: historyStore,
		PasswordChecker: passwordChecker,
		Housekeeper:     housekeeper,
	}
	totpStore := &totp.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	authenticatorTOTPConfig := authenticatorConfig.TOTP
	totpProvider := &totp.Provider{
		Store:  totpStore,
		Config: authenticatorTOTPConfig,
		Clock:  clock,
	}
	authenticatorOOBConfig := authenticatorConfig.OOB
	oobStore := &oob.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	oobProvider := &oob.Provider{
		Config: authenticatorOOBConfig,
		Store:  oobStore,
		Clock:  clock,
	}
	ratelimitLogger := ratelimit.NewLogger(factory)
	storageRedis := &ratelimit.StorageRedis{
		AppID: appID,
		Redis: handle,
	}
	limiter := &ratelimit.Limiter{
		Logger:  ratelimitLogger,
		Storage: storageRedis,
		Clock:   clock,
	}
	service3 := &service2.Service{
		Store:       store2,
		Password:    passwordProvider,
		TOTP:        totpProvider,
		OOBOTP:      oobProvider,
		RateLimiter: limiter,
	}
	verificationLogger := verification.NewLogger(factory)
	verificationConfig := appConfig.Verification
	verificationStoreRedis := &verification.StoreRedis{
		Redis: handle,
		AppID: appID,
		Clock: clock,
	}
	storePQ := &verification.StorePQ{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	verificationService := &verification.Service{
		Request:     request,
		Logger:      verificationLogger,
		Config:      verificationConfig,
		TrustProxy:  trustProxy,
		Clock:       clock,
		CodeStore:   verificationStoreRedis,
		ClaimStore:  storePQ,
		RateLimiter: limiter,
	}
	storeDeviceTokenRedis := &mfa.StoreDeviceTokenRedis{
		Redis: handle,
		AppID: appID,
		Clock: clock,
	}
	storeRecoveryCodePQ := &mfa.StoreRecoveryCodePQ{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	mfaService := &mfa.Service{
		DeviceTokens:  storeDeviceTokenRedis,
		RecoveryCodes: storeRecoveryCodePQ,
		Clock:         clock,
		Config:        authenticationConfig,
		RateLimiter:   limiter,
	}
	defaultTemplateLanguage := deps.ProvideDefaultTemplateLanguage(config)
	templateResolver := &template.Resolver{
		Resources:          manager,
		DefaultLanguageTag: defaultTemplateLanguage,
	}
	engine := &template.Engine{
		Resolver: templateResolver,
	}
	localizationConfig := appConfig.Localization
	staticAssetURLPrefix := environmentConfig.StaticAssetURLPrefix
	staticAssetResolver := &web.StaticAssetResolver{
		Context:            context,
		Config:             httpConfig,
		Localization:       localizationConfig,
		StaticAssetsPrefix: staticAssetURLPrefix,
		Resources:          manager,
	}
	translationService := &translation.Service{
		Context:           context,
		EnvironmentConfig: environmentConfig,
		TemplateEngine:    engine,
		StaticAssets:      staticAssetResolver,
	}
	welcomeMessageConfig := appConfig.WelcomeMessage
	queue := appProvider.TaskQueue
	welcomemessageProvider := &welcomemessage.Provider{
		Translation:          translationService,
		RateLimiter:          limiter,
		WelcomeMessageConfig: welcomeMessageConfig,
		TaskQueue:            queue,
	}
	rawCommands := &user.RawCommands{
		Store:                  store,
		Clock:                  clock,
		WelcomeMessageProvider: welcomemessageProvider,
	}
	idpsessionManager := &idpsession.Manager{
		Store:         storeRedis,
		Clock:         clock,
		Config:        sessionConfig,
		CookieFactory: cookieFactory,
		CookieDef:     cookieDef,
	}
	sessionManager := &oauth2.SessionManager{
		Store: grantStore,
		Clock: clock,
	}
	coordinator := &facade.Coordinator{
		Identities:      serviceService,
		Authenticators:  service3,
		Verification:    verificationService,
		MFA:             mfaService,
		Users:           rawCommands,
		PasswordHistory: historyStore,
		OAuth:           authorizationStore,
		IDPSessions:     idpsessionManager,
		OAuthSessions:   sessionManager,
		IdentityConfig:  identityConfig,
	}
	identityFacade := facade.IdentityFacade{
		Coordinator: coordinator,
	}
	queries := &user.Queries{
		Store:        store,
		Identities:   identityFacade,
		Verification: verificationService,
	}
	idTokenIssuer := &oidc.IDTokenIssuer{
		Secrets: oAuthKeyMaterials,
		BaseURL: endpointsProvider,
		Users:   queries,
		Clock:   clock,
	}
	accessTokenEncoding := &oauth2.AccessTokenEncoding{
		Secrets:    oAuthKeyMaterials,
		Clock:      clock,
		UserClaims: idTokenIssuer,
		BaseURL:    endpointsProvider,
	}
	oauthResolver := &oauth2.Resolver{
		TrustProxy:         trustProxy,
		Authorizations:     authorizationStore,
		AccessGrants:       grantStore,
		OfflineGrants:      grantStore,
		AccessTokenDecoder: accessTokenEncoding,
		Sessions:           provider,
		Clock:              clock,
	}
	sessionMiddleware := &session.Middleware{
		IDPSessionResolver:         resolver,
		AccessTokenSessionResolver: oauthResolver,
		AccessEvents:               eventProvider,
		Users:                      queries,
		Database:                   dbHandle,
	}
	return sessionMiddleware
}

var (
	_wireSystemClockValue = clock.NewSystemClock()
	_wireRandValue        = idpsession.Rand(rand.SecureRand)
)

func newSessionResolveHandler(p *deps.RequestProvider) http.Handler {
	appProvider := p.AppProvider
	config := appProvider.Config
	appConfig := config.AppConfig
	authenticationConfig := appConfig.Authentication
	identityConfig := appConfig.Identity
	secretConfig := config.SecretConfig
	databaseCredentials := deps.ProvideDatabaseCredentials(secretConfig)
	appID := appConfig.ID
	sqlBuilder := db.ProvideSQLBuilder(databaseCredentials, appID)
	request := p.Request
	context := deps.ProvideRequestContext(request)
	handle := appProvider.Database
	sqlExecutor := db.SQLExecutor{
		Context:  context,
		Database: handle,
	}
	store := &service.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	loginidStore := &loginid.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	loginIDConfig := identityConfig.LoginID
	manager := appProvider.Resources
	typeCheckerFactory := &loginid.TypeCheckerFactory{
		Config:    loginIDConfig,
		Resources: manager,
	}
	checker := &loginid.Checker{
		Config:             loginIDConfig,
		TypeCheckerFactory: typeCheckerFactory,
	}
	normalizerFactory := &loginid.NormalizerFactory{
		Config: loginIDConfig,
	}
	clockClock := _wireSystemClockValue
	provider := &loginid.Provider{
		Store:             loginidStore,
		Config:            loginIDConfig,
		Checker:           checker,
		NormalizerFactory: normalizerFactory,
		Clock:             clockClock,
	}
	oauthStore := &oauth.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	oauthProvider := &oauth.Provider{
		Store: oauthStore,
		Clock: clockClock,
	}
	anonymousStore := &anonymous.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	anonymousProvider := &anonymous.Provider{
		Store: anonymousStore,
		Clock: clockClock,
	}
	serviceService := &service.Service{
		Authentication: authenticationConfig,
		Identity:       identityConfig,
		Store:          store,
		LoginID:        provider,
		OAuth:          oauthProvider,
		Anonymous:      anonymousProvider,
	}
	factory := appProvider.LoggerFactory
	logger := verification.NewLogger(factory)
	verificationConfig := appConfig.Verification
	rootProvider := appProvider.RootProvider
	environmentConfig := rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	redisHandle := appProvider.Redis
	storeRedis := &verification.StoreRedis{
		Redis: redisHandle,
		AppID: appID,
		Clock: clockClock,
	}
	storePQ := &verification.StorePQ{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	ratelimitLogger := ratelimit.NewLogger(factory)
	storageRedis := &ratelimit.StorageRedis{
		AppID: appID,
		Redis: redisHandle,
	}
	limiter := &ratelimit.Limiter{
		Logger:  ratelimitLogger,
		Storage: storageRedis,
		Clock:   clockClock,
	}
	verificationService := &verification.Service{
		Request:     request,
		Logger:      logger,
		Config:      verificationConfig,
		TrustProxy:  trustProxy,
		Clock:       clockClock,
		CodeStore:   storeRedis,
		ClaimStore:  storePQ,
		RateLimiter: limiter,
	}
	resolveHandlerLogger := handler.NewResolveHandlerLogger(factory)
	resolveHandler := &handler.ResolveHandler{
		Identities:   serviceService,
		Verification: verificationService,
		Logger:       resolveHandlerLogger,
	}
	return resolveHandler
}
