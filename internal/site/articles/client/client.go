package client

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/alexfalkowski/go-service/env"
	se "github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/go-service/id"
	"github.com/alexfalkowski/go-service/net/http/content"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry/logger"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	th "github.com/alexfalkowski/go-service/transport/http"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"go.uber.org/fx"
)

var (
	// ErrNotFound when a resource is not there.
	ErrNotFound = errors.New("client: resource not found")

	// ErrInternal when a resource could not be processed.
	ErrInternal = errors.New("client: resource had internal issue")
)

// IsNotFound if the error is ErrNotFound.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
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
	Pool      *sync.BufferPool
	Content   *content.Content
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

	return &Client{pool: params.Pool, content: params.Content, client: cli}, nil
}

// Client for articles.
type Client struct {
	pool    *sync.BufferPool
	content *content.Content
	client  *http.Client
}

// Get the url and respond with res.
func (c *Client) Get(ctx context.Context, url string, res any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return se.Prefix("client: new request", err)
	}

	request.Header.Set(content.TypeKey, "application/yaml")

	response, err := c.client.Do(request)
	if err != nil {
		return se.Prefix("client: do", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		switch response.StatusCode {
		case http.StatusNotFound:
			return ErrNotFound
		default:
			return ErrInternal
		}
	}

	buffer := c.pool.Get()
	defer c.pool.Put(buffer)

	_, err = io.Copy(buffer, response.Body)
	if err != nil {
		return se.Prefix("client: copy", err)
	}

	media := c.content.NewFromMedia(response.Header.Get(content.TypeKey))

	if err := media.Encoder.Decode(buffer, res); err != nil {
		return se.Prefix("client: decode", err)
	}

	return nil
}
