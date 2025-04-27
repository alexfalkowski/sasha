package articles

import (
	"context"
	"net/http"

	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/go-service/net/http/mvc"
)

// Register books.
func Register(repo Repository) error {
	mvc.Route("GET /articles", func(_ context.Context) (mvc.View, *Model, error) {
		model := repo.GetArticles()

		return mvc.View("articles.tmpl"), model, nil
	})

	mvc.Route("GET /article/{slug}", func(ctx context.Context) (mvc.View, *Article, error) {
		slug := hc.Request(ctx).PathValue("slug")

		model := repo.GetArticle(slug)
		if model == nil {
			hc.Response(ctx).WriteHeader(http.StatusNotFound)
		}

		return mvc.View("article.tmpl"), model, nil
	})

	return nil
}
