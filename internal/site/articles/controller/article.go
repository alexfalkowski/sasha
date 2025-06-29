package controller

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
)

// NewArticle controller.
func NewArticle(repo repository.Repository, articleView, errorView *mvc.View) mvc.Controller[model.Article] {
	return func(ctx context.Context) (*mvc.View, *model.Article, error) {
		slug := meta.Request(ctx).PathValue("slug")

		model, err := repo.GetArticle(ctx, slug)
		if err != nil {
			return errorView, nil, err
		}

		return articleView, model, nil
	}
}
