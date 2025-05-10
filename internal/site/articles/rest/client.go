package rest

import (
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	th "github.com/alexfalkowski/go-service/transport/http"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"go.uber.org/fx"
)

// Params for rest.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	ID        id.Generator
	Config    *articles.Config
	Logger    *logger.Logger
	UserAgent env.UserAgent
}

// NewClient for rest.
func NewClient(params Params) *rest.Client {
	client, _ := th.NewClient(
		th.WithClientLogger(params.Logger), th.WithClientTracer(params.Tracer),
		th.WithClientMetrics(params.Meter), th.WithClientRetry(params.Config.Retry),
		th.WithClientUserAgent(params.UserAgent), th.WithClientTimeout(params.Config.Timeout),
		th.WithClientID(params.ID),
	)

	return rest.NewClient(
		rest.WithClientRoundTripper(client.Transport),
		rest.WithClientTimeout(params.Config.Timeout),
	)
}
