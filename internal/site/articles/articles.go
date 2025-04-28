package articles

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/meta"
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(repo Repository) error {
	mvc.Get("/articles", func(_ context.Context) (*mvc.View, *Model, error) {
		model := repo.GetArticles()

		return mvc.NewView("articles/articles.tmpl"), model, nil
	})

	mvc.Put("/articles", func(_ context.Context) (*mvc.View, *Model, error) {
		model := repo.GetArticles()

		return mvc.NewPartialView("articles/articles.tmpl"), model, nil
	})

	mvc.Get("/article/{slug}", func(ctx context.Context) (*mvc.View, *Article, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)
		slug := req.PathValue("slug")

		model := repo.GetArticle(slug)
		if model == nil {
			res.WriteHeader(http.StatusNotFound)
		}

		return mvc.NewView("articles/article.tmpl"), model, nil
	})

	mvc.Put("/article/{slug}", func(ctx context.Context) (*mvc.View, *Article, error) {
		req := meta.Request(ctx)
		res := meta.Response(ctx)
		slug := req.PathValue("slug")

		model := repo.GetArticle(slug)
		if model == nil {
			res.WriteHeader(http.StatusNotFound)
		}

		return mvc.NewPartialView("articles/article.tmpl"), model, nil
	})

	mvc.StaticPathValue("GET /images/{image}", "image", "images")

	return nil
}
