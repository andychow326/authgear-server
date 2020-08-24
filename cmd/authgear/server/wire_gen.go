// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/deps"
	"github.com/authgear/authgear-server/pkg/lib/infra/task/executor"
	"github.com/authgear/authgear-server/pkg/lib/infra/task/queue"
)

// Injectors from wire.go:

func newConfigSourceController(p *deps.RootProvider) *configsource.Controller {
	serverConfig := p.ServerConfig
	factory := p.LoggerFactory
	localFSLogger := configsource.NewLocalFSLogger(factory)
	localFS := &configsource.LocalFS{
		Logger:       localFSLogger,
		ServerConfig: serverConfig,
	}
	controller := configsource.NewController(serverConfig, localFS)
	return controller
}

func newInProcessQueue(p *deps.AppProvider, e *executor.InProcessExecutor) *queue.InProcessQueue {
	handle := p.Database
	config := p.Config
	captureTaskContext := deps.ProvideCaptureTaskContext(config)
	inProcessQueue := &queue.InProcessQueue{
		Database:       handle,
		CaptureContext: captureTaskContext,
		Executor:       e,
	}
	return inProcessQueue
}
