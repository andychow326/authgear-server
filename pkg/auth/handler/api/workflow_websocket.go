package api

import (
	"net/http"

	"github.com/iawaknahc/originmatcher"

	"github.com/authgear/authgear-server/pkg/lib/infra/redis/appredis"
	"github.com/authgear/authgear-server/pkg/util/httproute"
	"github.com/authgear/authgear-server/pkg/util/log"
	"github.com/authgear/authgear-server/pkg/util/pubsub"
)

func ConfigureWorkflowWebsocketRoute(route httproute.Route) httproute.Route {
	return route.
		WithMethods("OPTIONS", "GET").
		WithPathPattern("/api/v1/workflows/:workflowid/ws")
}

type WorkflowWebsocketEventStore interface {
	ChannelName(workflowID string) (string, error)
}

type WorkflowWebsocketOriginMatcher interface {
	PrepareOriginMatcher() (*originmatcher.T, error)
}

type WorkflowWebsocketHandler struct {
	Events        WorkflowWebsocketEventStore
	LoggerFactory *log.Factory
	RedisHandle   *appredis.Handle
	OriginMatcher WorkflowWebsocketOriginMatcher
}

func (h *WorkflowWebsocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matcher, err := h.OriginMatcher.PrepareOriginMatcher()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler := &pubsub.HTTPHandler{
		RedisHub:      h.RedisHandle,
		Delegate:      h,
		LoggerFactory: h.LoggerFactory,
		OriginMatcher: matcher,
	}

	handler.ServeHTTP(w, r)
}

func (h *WorkflowWebsocketHandler) Accept(r *http.Request) (string, error) {
	workflowID := httproute.GetParam(r, "workflowid")
	return h.Events.ChannelName(workflowID)
}

func (h *WorkflowWebsocketHandler) OnRedisSubscribe(r *http.Request) error {
	return nil
}
