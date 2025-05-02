package client

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/env"
	se "github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	th "github.com/alexfalkowski/go-service/transport/http"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"go.uber.org/fx"
)

// IsNotFound if we received a StatusNotFound.
func IsNotFound(err error) bool {
	return status.Code(err) == http.StatusNotFound
}

// Params for konfig.
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

// NewClient for articles.
func NewClient(params Params) (*Client, error) {
	cli, err := th.NewClient(
		th.WithClientLogger(params.Logger), th.WithClientTracer(params.Tracer),
		th.WithClientMetrics(params.Meter), th.WithClientRetry(params.Config.Retry),
		th.WithClientUserAgent(params.UserAgent), th.WithClientTimeout(params.Config.Timeout),
		th.WithClientTLS(params.Config.TLS), th.WithClientID(params.ID))
	if err != nil {
		return nil, se.Prefix("client: new http", err)
	}

	return &Client{client: rest.NewClient(rest.WithClientRoundTripper(cli.Transport), rest.WithClientTimeout(params.Config.Timeout))}, nil
}

// Client for articles.
type Client struct {
	client *rest.Client
}

// Get the url and respond with res.
func (c *Client) Get(ctx context.Context, url string, res any) error {
	opts := &rest.Options{
		ContentType: "application/yaml",
		Response:    res,
	}

	return c.client.Get(ctx, url, opts)
}
