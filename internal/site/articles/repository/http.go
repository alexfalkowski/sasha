package repository

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	se "github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/mime"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/net/http/status"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	sm "github.com/alexfalkowski/sasha/internal/site/meta"
)

// NewRepository for books.
func NewRepository(info *sm.Info, config *articles.Config, client *rest.Client) Repository {
	return &HTTPRepository{info: info, config: config, client: client}
}

// HTTPRepository uses a client to get from a site (public bucket).
type HTTPRepository struct {
	info   *sm.Info
	config *articles.Config
	client *rest.Client
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
			articles.Meta = meta.Strings(ctx, "")
			articles.Info = r.info

			return articles, nil
		}

		err := &model.Error{
			Meta: meta.Strings(ctx, ""),
			Info: r.info,
			Err:  se.Prefix("repository: get articles", err),
		}

		return nil, err
	}

	articles.Meta = meta.Strings(ctx, "")
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

	article.Meta = meta.Strings(ctx, "")
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
	buffer := &bytes.Buffer{}
	opts := &rest.Options{
		ContentType: mime.MarkdownMediaType,
		Response:    buffer,
	}
	url := fmt.Sprintf("%s/%s/article.md", r.config.Address, slug)

	if err := r.get(ctx, url, opts); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (r *HTTPRepository) get(ctx context.Context, url string, opts *rest.Options) error {
	if err := r.client.Get(ctx, url, opts); err != nil {
		if r.isNotFound(err) {
			err = ErrNotFound
		}

		err := &model.Error{
			Meta: meta.Strings(ctx, ""),
			Info: r.info,
			Err:  se.Prefix("repository: get url", err),
		}

		return err
	}

	return nil
}

func (r *HTTPRepository) isNotFound(err error) bool {
	return status.Code(err) == http.StatusNotFound
}
