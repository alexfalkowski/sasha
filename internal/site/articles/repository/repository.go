package repository

import (
	"context"
	"errors"

	"github.com/alexfalkowski/sasha/internal/site/articles/model"
)

// ErrNotFound when an article is not there.
var ErrNotFound = errors.New("not found")

// IsNotFound if the error is ErrNotFound.
//
//nolint:errorlint
func IsNotFound(err error) bool {
	if err, ok := err.(*model.Error); ok {
		return errors.Is(err.Err, ErrNotFound)
	}

	return errors.Is(err, ErrNotFound)
}

// Repository for books.
type Repository interface {
	// GetArticles from storage.
	GetArticles(ctx context.Context) (*model.Articles, error)

	// GetArticle by slug.
	GetArticle(ctx context.Context, slug string) (*model.Article, error)
}
