package repository

import (
	"context"

	"github.com/alexfalkowski/sasha/internal/site/articles/model"
)

// Repository for books.
type Repository interface {
	// GetArticles from storage.
	GetArticles(ctx context.Context) (*model.Articles, error)

	// GetArticle by slug.
	GetArticle(ctx context.Context, slug string) (*model.Article, error)
}
