package controller

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/meta"
	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
)

// NewArticlesController for articles.
func NewArticlesController(repo repository.Repository, articlesView, errorView *mvc.View) mvc.Controller[model.Articles] {
	return func(ctx context.Context) (*mvc.View, *model.Articles, error) {
		res := meta.Response(ctx)

		model, err := repo.GetArticles(ctx)
		if err != nil {
			if repository.IsNotFound(err) {
				res.WriteHeader(http.StatusNotFound)
			}

			return errorView, nil, err
		}

		return articlesView, model, nil
	}
}

// NewArticleController for articles.
func NewArticleController(repo repository.Repository, articleView, errorView *mvc.View) mvc.Controller[model.Article] {
	return func(ctx context.Context) (*mvc.View, *model.Article, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)
		slug := req.PathValue("slug")

		model, err := repo.GetArticle(ctx, slug)
		if err != nil {
			if repository.IsNotFound(err) {
				res.WriteHeader(http.StatusNotFound)
			}

			return errorView, nil, err
		}

		return articleView, model, nil
	}
}
