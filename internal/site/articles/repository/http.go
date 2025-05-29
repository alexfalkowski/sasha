package repository

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/mime"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	"github.com/alexfalkowski/go-service/v2/sync"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/rest"
	"github.com/alexfalkowski/sasha/internal/site/meta"
	"go.uber.org/fx"
)

// Params for repository.
type Params struct {
	fx.In

	Info   *meta.Info
	Config *articles.Config
	Client *rest.Client
	Pool   *sync.BufferPool
}

// NewRepository for repository.
func NewRepository(params Params) Repository {
	return &HTTPRepository{
		info:   params.Info,
		config: params.Config,
		client: params.Client,
		pool:   params.Pool,
	}
}

// HTTPRepository uses a client to get from a site (public bucket).
type HTTPRepository struct {
	info   *meta.Info
	config *articles.Config
	client *rest.Client
	pool   *sync.BufferPool
}

// GetArticles from the public bucket.
func (r *HTTPRepository) GetArticles(ctx context.Context) (*model.Articles, error) {
	articles := &model.Articles{}
	opts := &rest.Options{
		ContentType: mime.YAMLMediaType,
		Response:    articles,
	}
	url := r.config.Address + "/articles.yml"

	if err := r.client.Get(ctx, url, opts); err != nil {
		if r.isNotFound(err) {
			articles.Info = r.info

			return articles, nil
		}

		err := model.NewError(http.StatusInternalServerError, errors.Prefix("repository: get articles", err))
		err.Info = r.info

		return nil, err
	}

	articles.Info = r.info

	return articles, nil
}

// GetArticle by slug from the public bucket.
func (r *HTTPRepository) GetArticle(ctx context.Context, slug string) (*model.Article, error) {
	article, err := r.getArticleConfig(ctx, slug)
	if err != nil {
		return nil, err
	}

	body, err := r.getArticleBody(ctx, slug)
	if err != nil {
		return nil, err
	}

	article.Info = r.info
	article.Body = body

	return article, nil
}

func (r *HTTPRepository) getArticleConfig(ctx context.Context, slug string) (*model.Article, error) {
	article := &model.Article{}
	opts := &rest.Options{
		ContentType: mime.YAMLMediaType,
		Response:    article,
	}
	url := fmt.Sprintf("%s/%s/article.yml", r.config.Address, slug)

	if err := r.get(ctx, url, opts); err != nil {
		return nil, err
	}

	return article, nil
}

func (r *HTTPRepository) getArticleBody(ctx context.Context, slug string) ([]byte, error) {
	buffer := r.pool.Get()
	defer r.pool.Put(buffer)

	opts := &rest.Options{
		ContentType: mime.MarkdownMediaType,
		Response:    buffer,
	}
	url := fmt.Sprintf("%s/%s/article.md", r.config.Address, slug)

	if err := r.get(ctx, url, opts); err != nil {
		return nil, err
	}

	return r.pool.Copy(buffer), nil
}

func (r *HTTPRepository) get(ctx context.Context, url string, opts *rest.Options) error {
	if err := r.client.Get(ctx, url, opts); err != nil {
		var code int

		if r.isNotFound(err) {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}

		err := model.NewError(code, errors.Prefix("repository: get url", err))
		err.Info = r.info

		return err
	}

	return nil
}

func (r *HTTPRepository) isNotFound(err error) bool {
	return status.Code(err) == http.StatusNotFound
}
