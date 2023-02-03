// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package worker

import (
	"github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/elasticsearch"
	"github.com/authgear/authgear-server/pkg/lib/hook"
	"github.com/authgear/authgear-server/pkg/lib/infra/mail"
	"github.com/authgear/authgear-server/pkg/lib/infra/sms"
	"github.com/authgear/authgear-server/pkg/lib/infra/task"
	"github.com/authgear/authgear-server/pkg/lib/infra/task/executor"
	"github.com/authgear/authgear-server/pkg/worker/tasks"
)

// Injectors from wire.go:

func newInProcessExecutor(p *deps.RootProvider) *executor.InProcessExecutor {
	factory := p.LoggerFactory
	inProcessExecutorLogger := executor.NewInProcessExecutorLogger(factory)
	restoreTaskContext := deps.ProvideRestoreTaskContext(p)
	inProcessExecutor := &executor.InProcessExecutor{
		Logger:         inProcessExecutorLogger,
		RestoreContext: restoreTaskContext,
	}
	return inProcessExecutor
}

func newSendMessagesTask(p *deps.TaskProvider) task.Task {
	appProvider := p.AppProvider
	factory := appProvider.LoggerFactory
	logger := mail.NewLogger(factory)
	rootProvider := appProvider.RootProvider
	environmentConfig := rootProvider.EnvironmentConfig
	devMode := environmentConfig.DevMode
	appContext := appProvider.AppContext
	config := appContext.Config
	secretConfig := config.SecretConfig
	smtpServerCredentials := deps.ProvideSMTPServerCredentials(secretConfig)
	dialer := mail.NewGomailDialer(smtpServerCredentials)
	sender := &mail.Sender{
		Logger:       logger,
		DevMode:      devMode,
		GomailDialer: dialer,
	}
	smsLogger := sms.NewLogger(factory)
	appConfig := config.AppConfig
	messagingConfig := appConfig.Messaging
	twilioCredentials := deps.ProvideTwilioCredentials(secretConfig)
	twilioClient := sms.NewTwilioClient(twilioCredentials)
	nexmoCredentials := deps.ProvideNexmoCredentials(secretConfig)
	nexmoClient := sms.NewNexmoClient(nexmoCredentials)
	customSMSProviderConfig := deps.ProvideCustomSMSProviderConfig(secretConfig)
	context := p.Context
	denoEndpoint := environmentConfig.DenoEndpoint
	hookConfig := appConfig.Hook
	hookLogger := hook.NewLogger(factory)
	syncDenoClient := hook.NewSyncDenoClient(denoEndpoint, hookConfig, hookLogger)
	asyncDenoClient := hook.NewAsyncDenoClient(denoEndpoint, hookLogger)
	denoClientFactory := hook.NewDenoClientFactory(denoEndpoint, hookLogger)
	manager := appContext.Resources
	denoHookImpl := &hook.DenoHookImpl{
		Context:           context,
		SyncDenoClient:    syncDenoClient,
		AsyncDenoClient:   asyncDenoClient,
		DenoClientFactory: denoClientFactory,
		ResourceManager:   manager,
	}
	webhookKeyMaterials := deps.ProvideWebhookKeyMaterials(secretConfig)
	syncHTTPClient := hook.NewSyncHTTPClient(hookConfig)
	asyncHTTPClient := hook.NewAsyncHTTPClient()
	webHookImpl := &hook.WebHookImpl{
		Secret:    webhookKeyMaterials,
		SyncHTTP:  syncHTTPClient,
		AsyncHTTP: asyncHTTPClient,
	}
	customClient := sms.NewCustomClient(customSMSProviderConfig, denoHookImpl, webHookImpl)
	client := &sms.Client{
		Logger:          smsLogger,
		DevMode:         devMode,
		MessagingConfig: messagingConfig,
		TwilioClient:    twilioClient,
		NexmoClient:     nexmoClient,
		CustomClient:    customClient,
	}
	sendMessagesLogger := tasks.NewSendMessagesLogger(factory)
	sendMessagesTask := &tasks.SendMessagesTask{
		EmailSender: sender,
		SMSClient:   client,
		Logger:      sendMessagesLogger,
	}
	return sendMessagesTask
}

func newReindexUserTask(p *deps.TaskProvider) task.Task {
	appProvider := p.AppProvider
	appContext := appProvider.AppContext
	config := appContext.Config
	secretConfig := config.SecretConfig
	elasticsearchCredentials := deps.ProvideElasticsearchCredentials(secretConfig)
	client := elasticsearch.NewClient(elasticsearchCredentials)
	factory := appProvider.LoggerFactory
	reindexUserLogger := tasks.NewReindexUserLogger(factory)
	reindexUserTask := &tasks.ReindexUserTask{
		Client: client,
		Logger: reindexUserLogger,
	}
	return reindexUserTask
}
