package controller

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
)

// NewArticlesController for articles.
func NewArticlesController(repo repository.Repository, articlesView, errorView *mvc.View) mvc.Controller[model.Articles] {
	return func(ctx context.Context) (*mvc.View, *model.Articles, error) {
		model, err := repo.GetArticles(ctx)
		if err != nil {
			return errorView, nil, err
		}

		return articlesView, model, nil
	}
}
