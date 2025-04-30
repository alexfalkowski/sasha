package articles

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/meta"
	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
)

// Register books.
func Register(repo repository.Repository) error {
	errorView, errorPartialView := mvc.NewViewPair("articles/error.tmpl")
	articlesView, articlesPartialView := mvc.NewViewPair("articles/articles.tmpl")
	articleView, articlePartialView := mvc.NewViewPair("articles/article.tmpl")

	mvc.Get("/articles", func(ctx context.Context) (*mvc.View, *model.Articles, error) {
		model, err := repo.GetArticles(ctx)
		if err != nil {
			return errorView, nil, err
		}

		return articlesView, model, nil
	})

	mvc.Put("/articles", func(ctx context.Context) (*mvc.View, *model.Articles, error) {
		model, err := repo.GetArticles(ctx)
		if err != nil {
			return errorPartialView, nil, err
		}

		return articlesPartialView, model, err
	})

	mvc.Get("/article/{slug}", func(ctx context.Context) (*mvc.View, *model.Article, error) {
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
	})

	mvc.Put("/article/{slug}", func(ctx context.Context) (*mvc.View, *model.Article, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)
		slug := req.PathValue("slug")

		model, err := repo.GetArticle(ctx, slug)
		if err != nil {
			if repository.IsNotFound(err) {
				res.WriteHeader(http.StatusNotFound)
			}

			return errorPartialView, nil, err
		}

		return articlePartialView, model, err
	})

	return nil
}
