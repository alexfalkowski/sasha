package articles

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/meta"
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(repo Repository) error {
	errorView, errorPartialView := mvc.NewViewPair("articles/error.tmpl")
	articlesView, articlesPartialView := mvc.NewViewPair("articles/articles.tmpl")
	articleView, articlePartialView := mvc.NewViewPair("articles/article.tmpl")

	mvc.Get("/articles", func(ctx context.Context) (*mvc.View, *Articles, error) {
		model, err := repo.GetArticles(ctx)
		if err != nil {
			return errorView, nil, err
		}

		return articlesView, model, nil
	})

	mvc.Put("/articles", func(ctx context.Context) (*mvc.View, *Articles, error) {
		model, err := repo.GetArticles(ctx)
		if err != nil {
			return errorPartialView, nil, err
		}

		return articlesPartialView, model, err
	})

	mvc.Get("/article/{slug}", func(ctx context.Context) (*mvc.View, *Article, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)
		slug := req.PathValue("slug")

		model, err := repo.GetArticle(ctx, slug)
		if err != nil {
			return errorView, nil, err
		}

		if model == nil {
			res.WriteHeader(http.StatusNotFound)
		}

		return articleView, model, nil
	})

	mvc.Put("/article/{slug}", func(ctx context.Context) (*mvc.View, *Article, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)
		slug := req.PathValue("slug")

		model, err := repo.GetArticle(ctx, slug)
		if err != nil {
			return errorPartialView, nil, err
		}

		if model == nil {
			res.WriteHeader(http.StatusNotFound)
		}

		return articlePartialView, model, err
	})

	return nil
}
