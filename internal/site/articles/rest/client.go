package rest

import (
	"context"
	"fmt"
	"time"

	"github.com/alexfalkowski/go-service/cache/cacheable"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/go-service/transport/http"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"go.uber.org/fx"
)

// Params for rest.
type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	ID        id.Generator
	Cache     cacheable.Interface
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	Config    *articles.Config
	Logger    *logger.Logger
	UserAgent env.UserAgent
}

// NewClient for rest.
func NewClient(params Params) *Client {
	client, _ := http.NewClient(
		http.WithClientLogger(params.Logger), http.WithClientTracer(params.Tracer),
		http.WithClientMetrics(params.Meter), http.WithClientRetry(params.Config.Retry),
		http.WithClientUserAgent(params.UserAgent), http.WithClientTimeout(params.Config.Timeout),
		http.WithClientID(params.ID),
	)

	return &Client{
		client: rest.NewClient(
			rest.WithClientRoundTripper(client.Transport),
			rest.WithClientTimeout(params.Config.Timeout),
		),
		cache: params.Cache,
	}
}

// Options is an alias of rest.Options.
type Options = rest.Options

// Client for articles.
type Client struct {
	client *rest.Client
	cache  cacheable.Interface
}

// Get a url with opts. This uses a cache for 1 hour.
func (c *Client) Get(ctx context.Context, url string, opts *Options) error {
	key := fmt.Sprintf("%s/%s", opts.ContentType, url)

	ok, err := c.cache.Get(ctx, key, opts.Response)
	if err != nil {
		return err
	}

	if ok {
		return nil
	}

	if err := c.client.Get(ctx, url, opts); err != nil {
		return err
	}

	return c.cache.Persist(ctx, key, opts.Response, time.Hour)
}
