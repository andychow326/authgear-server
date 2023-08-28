// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package portal

import (
	"github.com/authgear/authgear-server/pkg/lib/admin/authz"
	"github.com/authgear/authgear-server/pkg/lib/analytic"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/hook"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/auditdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/lib/infra/mail"
	"github.com/authgear/authgear-server/pkg/lib/infra/middleware"
	"github.com/authgear/authgear-server/pkg/lib/tutorial"
	"github.com/authgear/authgear-server/pkg/lib/usage"
	"github.com/authgear/authgear-server/pkg/portal/appresource/factory"
	"github.com/authgear/authgear-server/pkg/portal/appsecret"
	"github.com/authgear/authgear-server/pkg/portal/deps"
	"github.com/authgear/authgear-server/pkg/portal/endpoint"
	"github.com/authgear/authgear-server/pkg/portal/graphql"
	"github.com/authgear/authgear-server/pkg/portal/lib/plan"
	"github.com/authgear/authgear-server/pkg/portal/libstripe"
	"github.com/authgear/authgear-server/pkg/portal/loader"
	"github.com/authgear/authgear-server/pkg/portal/service"
	"github.com/authgear/authgear-server/pkg/portal/session"
	"github.com/authgear/authgear-server/pkg/portal/smtp"
	"github.com/authgear/authgear-server/pkg/portal/task"
	"github.com/authgear/authgear-server/pkg/portal/task/tasks"
	"github.com/authgear/authgear-server/pkg/portal/transport"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/intl"
	"github.com/authgear/authgear-server/pkg/util/template"
	"net/http"
)

import (
	_ "github.com/authgear/authgear-server/pkg/auth"
)

// Injectors from wire.go:

func newPanicMiddleware(p *deps.RequestProvider) httproute.Middleware {
	rootProvider := p.RootProvider
	factory := rootProvider.LoggerFactory
	panicMiddlewareLogger := middleware.NewPanicMiddlewareLogger(factory)
	panicMiddleware := &middleware.PanicMiddleware{
		Logger: panicMiddlewareLogger,
	}
	return panicMiddleware
}

func newBodyLimitMiddleware(p *deps.RequestProvider) httproute.Middleware {
	bodyLimitMiddleware := &middleware.BodyLimitMiddleware{}
	return bodyLimitMiddleware
}

func newSentryMiddleware(p *deps.RequestProvider) httproute.Middleware {
	rootProvider := p.RootProvider
	hub := rootProvider.SentryHub
	environmentConfig := rootProvider.EnvironmentConfig
	trustProxy := environmentConfig.TrustProxy
	sentryMiddleware := &middleware.SentryMiddleware{
		SentryHub:  hub,
		TrustProxy: trustProxy,
	}
	return sentryMiddleware
}

func newSessionInfoMiddleware(p *deps.RequestProvider) httproute.Middleware {
	sessionInfoMiddleware := &session.SessionInfoMiddleware{}
	return sessionInfoMiddleware
}

func newSessionRequiredMiddleware(p *deps.RequestProvider) httproute.Middleware {
	sessionRequiredMiddleware := &session.SessionRequiredMiddleware{}
	return sessionRequiredMiddleware
}

func newGraphQLHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	environmentConfig := rootProvider.EnvironmentConfig
	devMode := environmentConfig.DevMode
	request := p.Request
	logFactory := rootProvider.LoggerFactory
	logger := graphql.NewLogger(logFactory)
	authgearConfig := rootProvider.AuthgearConfig
	adminAPIConfig := rootProvider.AdminAPIConfig
	controller := rootProvider.ConfigSourceController
	configSource := deps.ProvideConfigSource(controller)
	clock := _wireSystemClockValue
	adder := &authz.Adder{
		Clock: clock,
	}
	appHostSuffixes := environmentConfig.AppHostSuffixes
	appConfig := rootProvider.AppConfig
	context := deps.ProvideRequestContext(request)
	configServiceLogger := service.NewConfigServiceLogger(logFactory)
	domainImplementationType := rootProvider.DomainImplementation
	kubernetesConfig := rootProvider.KubernetesConfig
	kubernetesLogger := service.NewKubernetesLogger(logFactory)
	kubernetes := &service.Kubernetes{
		KubernetesConfig: kubernetesConfig,
		AppConfig:        appConfig,
		Logger:           kubernetesLogger,
	}
	configService := &service.ConfigService{
		Context:              context,
		Logger:               configServiceLogger,
		AppConfig:            appConfig,
		Controller:           controller,
		ConfigSource:         configSource,
		DomainImplementation: domainImplementationType,
		Kubernetes:           kubernetes,
	}
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	pool := rootProvider.Database
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(context, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, logFactory)
	sqlExecutor := globaldb.NewSQLExecutor(context, handle)
	domainService := &service.DomainService{
		Context:      context,
		Clock:        clock,
		DomainConfig: configService,
		SQLBuilder:   sqlBuilder,
		SQLExecutor:  sqlExecutor,
	}
	defaultDomainService := &service.DefaultDomainService{
		AppHostSuffixes: appHostSuffixes,
		AppConfig:       appConfig,
		Domains:         domainService,
	}
	adminAPIService := &service.AdminAPIService{
		AuthgearConfig: authgearConfig,
		AdminAPIConfig: adminAPIConfig,
		ConfigSource:   configSource,
		AuthzAdder:     adder,
		DefaultDomains: defaultDomainService,
	}
	appServiceLogger := service.NewAppServiceLogger(logFactory)
	mailConfig := rootProvider.MailConfig
	inProcessExecutorLogger := task.NewInProcessExecutorLogger(logFactory)
	mailLogger := mail.NewLogger(logFactory)
	smtpConfig := rootProvider.SMTPConfig
	smtpServerCredentials := deps.ProvideSMTPServerCredentials(smtpConfig)
	dialer := mail.NewGomailDialer(smtpServerCredentials)
	testModeEmailSuppressed := deps.ProvideTestModeEmailSuppressed()
	sender := &mail.Sender{
		Logger:                  mailLogger,
		DevMode:                 devMode,
		GomailDialer:            dialer,
		TestModeEmailSuppressed: testModeEmailSuppressed,
	}
	sendMessagesLogger := tasks.NewSendMessagesLogger(logFactory)
	sendMessagesTask := &tasks.SendMessagesTask{
		EmailSender: sender,
		Logger:      sendMessagesLogger,
	}
	inProcessExecutor := task.NewExecutor(inProcessExecutorLogger, sendMessagesTask)
	inProcessQueue := &task.InProcessQueue{
		Executor: inProcessExecutor,
	}
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	requestOriginProvider := &endpoint.RequestOriginProvider{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	endpointsProvider := &endpoint.EndpointsProvider{
		OriginProvider: requestOriginProvider,
	}
	manager := rootProvider.Resources
	defaultLanguageTag := _wireDefaultLanguageTagValue
	supportedLanguageTags := _wireSupportedLanguageTagsValue
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	collaboratorService := &service.CollaboratorService{
		Context:        context,
		Clock:          clock,
		SQLBuilder:     sqlBuilder,
		SQLExecutor:    sqlExecutor,
		MailConfig:     mailConfig,
		TaskQueue:      inProcessQueue,
		Endpoints:      endpointsProvider,
		TemplateEngine: engine,
		AdminAPI:       adminAPIService,
		AppConfigs:     configService,
	}
	authzService := &service.AuthzService{
		Context:       context,
		Configs:       configService,
		Collaborators: collaboratorService,
	}
	appBaseResources := deps.ProvideAppBaseResources(rootProvider)
	storeImpl := &tutorial.StoreImpl{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	tutorialService := &tutorial.Service{
		Store: storeImpl,
	}
	denoEndpoint := environmentConfig.DenoEndpoint
	hookLogger := hook.NewLogger(logFactory)
	denoClientImpl := ProvideDenoClient(denoEndpoint, hookLogger)
	managerFactory := &factory.ManagerFactory{
		Context:          context,
		AppBaseResources: appBaseResources,
		Tutorials:        tutorialService,
		DenoClient:       denoClientImpl,
		Clock:            clock,
	}
	store := &plan.Store{
		Clock:       clock,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	planService := &plan.Service{
		PlanStore: store,
		AppConfig: appConfig,
	}
	globalredisHandle := rootProvider.GlobalRedisHandle
	appSecretVisitTokenStoreImpl := &appsecret.AppSecretVisitTokenStoreImpl{
		Context: context,
		Redis:   globalredisHandle,
	}
	appService := &service.AppService{
		Logger:                   appServiceLogger,
		SQLBuilder:               sqlBuilder,
		SQLExecutor:              sqlExecutor,
		AppConfig:                appConfig,
		AppConfigs:               configService,
		AppAuthz:                 authzService,
		DefaultDomains:           defaultDomainService,
		Resources:                manager,
		AppResMgrFactory:         managerFactory,
		Plan:                     planService,
		Clock:                    clock,
		AppSecretVisitTokenStore: appSecretVisitTokenStoreImpl,
	}
	userLoader := loader.NewUserLoader(adminAPIService, appService, collaboratorService)
	appLoader := loader.NewAppLoader(appService, authzService)
	domainLoader := loader.NewDomainLoader(domainService, authzService)
	collaboratorLoader := loader.NewCollaboratorLoader(collaboratorService, authzService)
	collaboratorInvitationLoader := loader.NewCollaboratorInvitationLoader(collaboratorService, authzService)
	smtpService := &smtp.Service{
		Context: context,
	}
	auditDatabaseCredentials := deps.ProvideAuditDatabaseCredentials(environmentConfig)
	readHandle := auditdb.NewReadHandle(context, pool, databaseEnvironmentConfig, auditDatabaseCredentials, logFactory)
	auditdbSQLBuilder := auditdb.NewSQLBuilder(auditDatabaseCredentials)
	readSQLExecutor := auditdb.NewReadSQLExecutor(context, readHandle)
	auditDBReadStore := &analytic.AuditDBReadStore{
		SQLBuilder:  auditdbSQLBuilder,
		SQLExecutor: readSQLExecutor,
	}
	analyticConfig := rootProvider.AnalyticConfig
	chartService := &analytic.ChartService{
		Database:       readHandle,
		AuditStore:     auditDBReadStore,
		Clock:          clock,
		AnalyticConfig: analyticConfig,
	}
	stripeConfig := rootProvider.StripeConfig
	libstripeLogger := libstripe.NewLogger(logFactory)
	api := libstripe.NewClientAPI(stripeConfig, libstripeLogger)
	stripeCache := libstripe.NewStripeCache()
	libstripeService := &libstripe.Service{
		ClientAPI:         api,
		Logger:            libstripeLogger,
		Context:           context,
		Plans:             planService,
		GlobalRedisHandle: globalredisHandle,
		Cache:             stripeCache,
		Clock:             clock,
		StripeConfig:      stripeConfig,
		Endpoints:         endpointsProvider,
	}
	configsourceStore := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	globalDBStore := &usage.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	subscriptionService := &service.SubscriptionService{
		SQLBuilder:        sqlBuilder,
		SQLExecutor:       sqlExecutor,
		ConfigSourceStore: configsourceStore,
		PlanStore:         store,
		UsageStore:        globalDBStore,
		Clock:             clock,
		AppConfig:         appConfig,
	}
	nftIndexerAPIEndpoint := environmentConfig.NFTIndexerAPIEndpoint
	nftService := &service.NFTService{
		APIEndpoint: nftIndexerAPIEndpoint,
	}
	remoteIP := deps.ProvideRemoteIP(request, trustProxy)
	userAgentString := deps.ProvideUserAgentString(request)
	writeHandle := auditdb.NewWriteHandle(context, pool, databaseEnvironmentConfig, auditDatabaseCredentials, logFactory)
	auditService := &service.AuditService{
		Context:         context,
		RemoteIP:        remoteIP,
		UserAgentString: userAgentString,
		Request:         request,
		Apps:            appService,
		Authgear:        authgearConfig,
		DenoEndpoint:    denoEndpoint,
		SQLBuilder:      sqlBuilder,
		SQLExecutor:     sqlExecutor,
		AuditDatabase:   writeHandle,
		Clock:           clock,
		LoggerFactory:   logFactory,
	}
	graphqlContext := &graphql.Context{
		Request:                 request,
		GQLLogger:               logger,
		Users:                   userLoader,
		Apps:                    appLoader,
		Domains:                 domainLoader,
		Collaborators:           collaboratorLoader,
		CollaboratorInvitations: collaboratorInvitationLoader,
		AuthzService:            authzService,
		AppService:              appService,
		DomainService:           domainService,
		CollaboratorService:     collaboratorService,
		SMTPService:             smtpService,
		AppResMgrFactory:        managerFactory,
		AnalyticChartService:    chartService,
		TutorialService:         tutorialService,
		StripeService:           libstripeService,
		SubscriptionService:     subscriptionService,
		NFTService:              nftService,
		DenoService:             denoClientImpl,
		AuditService:            auditService,
	}
	graphQLHandler := &transport.GraphQLHandler{
		DevMode:        devMode,
		GraphQLContext: graphqlContext,
		Database:       handle,
		AuditDatabase:  readHandle,
	}
	return graphQLHandler
}

var (
	_wireSystemClockValue           = clock.NewSystemClock()
	_wireDefaultLanguageTagValue    = template.DefaultLanguageTag(intl.BuiltinBaseLanguage)
	_wireSupportedLanguageTagsValue = template.SupportedLanguageTags([]string{intl.BuiltinBaseLanguage})
)

func newSystemConfigHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	authgearConfig := rootProvider.AuthgearConfig
	appConfig := rootProvider.AppConfig
	searchConfig := rootProvider.SearchConfig
	web3Config := rootProvider.Web3Config
	auditLogConfig := rootProvider.AuditLogConfig
	analyticConfig := rootProvider.AnalyticConfig
	googleTagManagerConfig := rootProvider.GoogleTagManagerConfig
	portalFrontendSentryConfig := rootProvider.PortalFrontendSentryConfig
	manager := rootProvider.Resources
	systemConfigProvider := &service.SystemConfigProvider{
		AuthgearConfig:       authgearConfig,
		AppConfig:            appConfig,
		SearchConfig:         searchConfig,
		Web3Config:           web3Config,
		AuditLogConfig:       auditLogConfig,
		AnalyticConfig:       analyticConfig,
		GTMConfig:            googleTagManagerConfig,
		FrontendSentryConfig: portalFrontendSentryConfig,
		Resources:            manager,
	}
	filesystemCache := rootProvider.FilesystemCache
	systemConfigHandler := &transport.SystemConfigHandler{
		SystemConfig:    systemConfigProvider,
		FilesystemCache: filesystemCache,
	}
	return systemConfigHandler
}

func newAdminAPIHandler(p *deps.RequestProvider) http.Handler {
	request := p.Request
	context := deps.ProvideRequestContext(request)
	rootProvider := p.RootProvider
	pool := rootProvider.Database
	environmentConfig := rootProvider.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	logFactory := rootProvider.LoggerFactory
	handle := globaldb.NewHandle(context, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, logFactory)
	configServiceLogger := service.NewConfigServiceLogger(logFactory)
	appConfig := rootProvider.AppConfig
	controller := rootProvider.ConfigSourceController
	configSource := deps.ProvideConfigSource(controller)
	domainImplementationType := rootProvider.DomainImplementation
	kubernetesConfig := rootProvider.KubernetesConfig
	kubernetesLogger := service.NewKubernetesLogger(logFactory)
	kubernetes := &service.Kubernetes{
		KubernetesConfig: kubernetesConfig,
		AppConfig:        appConfig,
		Logger:           kubernetesLogger,
	}
	configService := &service.ConfigService{
		Context:              context,
		Logger:               configServiceLogger,
		AppConfig:            appConfig,
		Controller:           controller,
		ConfigSource:         configSource,
		DomainImplementation: domainImplementationType,
		Kubernetes:           kubernetes,
	}
	clockClock := _wireSystemClockValue
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	sqlExecutor := globaldb.NewSQLExecutor(context, handle)
	mailConfig := rootProvider.MailConfig
	inProcessExecutorLogger := task.NewInProcessExecutorLogger(logFactory)
	logger := mail.NewLogger(logFactory)
	devMode := environmentConfig.DevMode
	smtpConfig := rootProvider.SMTPConfig
	smtpServerCredentials := deps.ProvideSMTPServerCredentials(smtpConfig)
	dialer := mail.NewGomailDialer(smtpServerCredentials)
	testModeEmailSuppressed := deps.ProvideTestModeEmailSuppressed()
	sender := &mail.Sender{
		Logger:                  logger,
		DevMode:                 devMode,
		GomailDialer:            dialer,
		TestModeEmailSuppressed: testModeEmailSuppressed,
	}
	sendMessagesLogger := tasks.NewSendMessagesLogger(logFactory)
	sendMessagesTask := &tasks.SendMessagesTask{
		EmailSender: sender,
		Logger:      sendMessagesLogger,
	}
	inProcessExecutor := task.NewExecutor(inProcessExecutorLogger, sendMessagesTask)
	inProcessQueue := &task.InProcessQueue{
		Executor: inProcessExecutor,
	}
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	requestOriginProvider := &endpoint.RequestOriginProvider{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	endpointsProvider := &endpoint.EndpointsProvider{
		OriginProvider: requestOriginProvider,
	}
	manager := rootProvider.Resources
	defaultLanguageTag := _wireDefaultLanguageTagValue
	supportedLanguageTags := _wireSupportedLanguageTagsValue
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	authgearConfig := rootProvider.AuthgearConfig
	adminAPIConfig := rootProvider.AdminAPIConfig
	adder := &authz.Adder{
		Clock: clockClock,
	}
	appHostSuffixes := environmentConfig.AppHostSuffixes
	domainService := &service.DomainService{
		Context:      context,
		Clock:        clockClock,
		DomainConfig: configService,
		SQLBuilder:   sqlBuilder,
		SQLExecutor:  sqlExecutor,
	}
	defaultDomainService := &service.DefaultDomainService{
		AppHostSuffixes: appHostSuffixes,
		AppConfig:       appConfig,
		Domains:         domainService,
	}
	adminAPIService := &service.AdminAPIService{
		AuthgearConfig: authgearConfig,
		AdminAPIConfig: adminAPIConfig,
		ConfigSource:   configSource,
		AuthzAdder:     adder,
		DefaultDomains: defaultDomainService,
	}
	collaboratorService := &service.CollaboratorService{
		Context:        context,
		Clock:          clockClock,
		SQLBuilder:     sqlBuilder,
		SQLExecutor:    sqlExecutor,
		MailConfig:     mailConfig,
		TaskQueue:      inProcessQueue,
		Endpoints:      endpointsProvider,
		TemplateEngine: engine,
		AdminAPI:       adminAPIService,
		AppConfigs:     configService,
	}
	authzService := &service.AuthzService{
		Context:       context,
		Configs:       configService,
		Collaborators: collaboratorService,
	}
	adminAPILogger := transport.NewAdminAPILogger(logFactory)
	adminAPIHandler := &transport.AdminAPIHandler{
		Database: handle,
		Authz:    authzService,
		AdminAPI: adminAPIService,
		Logger:   adminAPILogger,
	}
	return adminAPIHandler
}

func newStaticAssetsHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	manager := rootProvider.Resources
	staticAssetsHandler := &transport.StaticAssetsHandler{
		Resources: manager,
	}
	return staticAssetsHandler
}

func newStripeWebhookHandler(p *deps.RequestProvider) http.Handler {
	rootProvider := p.RootProvider
	stripeConfig := rootProvider.StripeConfig
	logFactory := rootProvider.LoggerFactory
	logger := libstripe.NewLogger(logFactory)
	api := libstripe.NewClientAPI(stripeConfig, logger)
	request := p.Request
	context := deps.ProvideRequestContext(request)
	clockClock := _wireSystemClockValue
	environmentConfig := rootProvider.EnvironmentConfig
	globalDatabaseCredentialsEnvironmentConfig := &environmentConfig.GlobalDatabase
	sqlBuilder := globaldb.NewSQLBuilder(globalDatabaseCredentialsEnvironmentConfig)
	pool := rootProvider.Database
	databaseEnvironmentConfig := &environmentConfig.DatabaseConfig
	handle := globaldb.NewHandle(context, pool, globalDatabaseCredentialsEnvironmentConfig, databaseEnvironmentConfig, logFactory)
	sqlExecutor := globaldb.NewSQLExecutor(context, handle)
	store := &plan.Store{
		Clock:       clockClock,
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	appConfig := rootProvider.AppConfig
	planService := &plan.Service{
		PlanStore: store,
		AppConfig: appConfig,
	}
	globalredisHandle := rootProvider.GlobalRedisHandle
	stripeCache := libstripe.NewStripeCache()
	trustProxy := environmentConfig.TrustProxy
	httpHost := deps.ProvideHTTPHost(request, trustProxy)
	httpProto := deps.ProvideHTTPProto(request, trustProxy)
	requestOriginProvider := &endpoint.RequestOriginProvider{
		HTTPHost:  httpHost,
		HTTPProto: httpProto,
	}
	endpointsProvider := &endpoint.EndpointsProvider{
		OriginProvider: requestOriginProvider,
	}
	libstripeService := &libstripe.Service{
		ClientAPI:         api,
		Logger:            logger,
		Context:           context,
		Plans:             planService,
		GlobalRedisHandle: globalredisHandle,
		Cache:             stripeCache,
		Clock:             clockClock,
		StripeConfig:      stripeConfig,
		Endpoints:         endpointsProvider,
	}
	stripeWebhookLogger := transport.NewStripeWebhookLogger(logFactory)
	configsourceStore := &configsource.Store{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	globalDBStore := &usage.GlobalDBStore{
		SQLBuilder:  sqlBuilder,
		SQLExecutor: sqlExecutor,
	}
	subscriptionService := &service.SubscriptionService{
		SQLBuilder:        sqlBuilder,
		SQLExecutor:       sqlExecutor,
		ConfigSourceStore: configsourceStore,
		PlanStore:         store,
		UsageStore:        globalDBStore,
		Clock:             clockClock,
		AppConfig:         appConfig,
	}
	stripeWebhookHandler := &transport.StripeWebhookHandler{
		StripeService: libstripeService,
		Logger:        stripeWebhookLogger,
		Subscriptions: subscriptionService,
		Database:      handle,
	}
	return stripeWebhookHandler
}
