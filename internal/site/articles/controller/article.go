package controller

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/meta"
	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
)

// NewArticle controller.
func NewArticle(repo repository.Repository, articleView, errorView *mvc.View) mvc.Controller[model.Article] {
	return func(ctx context.Context) (*mvc.View, *model.Article, error) {
		slug := meta.Request(ctx).PathValue("slug")

		model, err := repo.GetArticle(ctx, slug)
		if err != nil {
			if repository.IsNotFound(err) {
				err = status.FromError(http.StatusNotFound, err)
			}

			return errorView, nil, err
		}

		return articleView, model, nil
	}
}
