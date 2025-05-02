package controller

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
)

// NewArticlesController for articles.
func NewArticlesController(repo repository.Repository, articlesView, errorView *mvc.View) mvc.Controller[model.Articles] {
	return func(ctx context.Context) (*mvc.View, *model.Articles, error) {
		model, err := repo.GetArticles(ctx)
		if err != nil {
			if repository.IsNotFound(err) {
				err = status.FromError(http.StatusNotFound, err)
			}

			return errorView, nil, err
		}

		return articlesView, model, nil
	}
}
