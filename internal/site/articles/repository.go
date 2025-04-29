package articles

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/sasha/internal/site/meta"
)

// Repository for books.
type Repository interface {
	// GetArticles from storage.
	GetArticles(ctx context.Context) (*Articles, error)

	// GetArticle by slug.
	GetArticle(ctx context.Context, slug string) (*Article, error)
}

// NewRepository for books.
func NewRepository(info *meta.Info, config *Config, client *Client) Repository {
	return &FileSystemRepository{info: info, config: config, client: client}
}

// FSRepository has books in a file.
type FileSystemRepository struct {
	info   *meta.Info
	config *Config
	client *Client
}

// GetArticles from a file.
func (r *FileSystemRepository) GetArticles(ctx context.Context) (*Articles, error) {
	site := &Articles{}

	if err := r.client.Get(ctx, r.config.Address+"/articles.yml", site); err != nil {
		e := &Error{
			err:  errors.Prefix("repository: get articles", err),
			Info: r.info,
		}

		return nil, e
	}

	site.Info = r.info

	return site, nil
}

// GetArticle by slug.
func (r *FileSystemRepository) GetArticle(ctx context.Context, slug string) (*Article, error) {
	article := &Article{}
	url := fmt.Sprintf("%s/%s/article.yml", r.config.Address, slug)

	if err := r.client.Get(ctx, url, article); err != nil {
		if r.client.IsNotFound(err) {
			return nil, nil
		}

		e := &Error{
			err:  errors.Prefix("repository: get article", err),
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
