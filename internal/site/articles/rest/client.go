package rest

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/v2/cache/cacher"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/id"
	"github.com/alexfalkowski/go-service/v2/net/http/rest"
	"github.com/alexfalkowski/go-service/v2/telemetry/logger"
	"github.com/alexfalkowski/go-service/v2/telemetry/metrics"
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/alexfalkowski/go-service/v2/transport/http"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"go.uber.org/fx"
)

// Params for rest.
type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	ID        id.Generator
	Cache     cacher.Cache
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
	cache  cacher.Cache
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

	return c.persist(ctx, key, opts.Response)
}

func (c *Client) persist(ctx context.Context, key string, response any) error {
	if err := c.cache.Persist(ctx, key, response, 15*time.Minute); err != nil {
		return err
	}

	_, err := c.cache.Get(ctx, key, response)

	return err
}
