package repository

import (
	"context"
	"fmt"

	se "github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/sasha/internal/site/articles/client"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/meta"
)

// NewRepository for books.
func NewRepository(info *meta.Info, config *articles.Config, client *client.Client) Repository {
	return &HTTPRepository{info: info, config: config, client: client}
}

// HTTPRepository uses a client to get from a site (public bucket).
type HTTPRepository struct {
	info   *meta.Info
	config *articles.Config
	client *client.Client
}

// GetArticles from the public bucket.
func (r *HTTPRepository) GetArticles(ctx context.Context) (*model.Articles, error) {
	site := &model.Articles{}

	if err := r.client.Get(ctx, r.config.Address+"/articles.yml", site); err != nil {
		if client.IsNotFound(err) {
			err = ErrNotFound
		}

		e := &model.Error{
			Err:  se.Prefix("repository: get articles", err),
			Info: r.info,
		}

		return nil, e
	}

	site.Info = r.info

	return site, nil
}

// GetArticle by slug from the public bucket.
func (r *HTTPRepository) GetArticle(ctx context.Context, slug string) (*model.Article, error) {
	article := &model.Article{}
	url := fmt.Sprintf("%s/%s/article.yml", r.config.Address, slug)

	if err := r.client.Get(ctx, url, article); err != nil {
		if client.IsNotFound(err) {
			err = ErrNotFound
		}

		e := &model.Error{
			Err:  se.Prefix("repository: get article", err),
			Info: r.info,
		}

		return nil, e
	}

	article.Info = r.info

	// Transform the images to URLs for the view,
	for _, image := range article.Images {
		image.Name = fmt.Sprintf("%s/%s/images/%s", r.config.Address, slug, image.Name)
	}

	return article, nil
}
