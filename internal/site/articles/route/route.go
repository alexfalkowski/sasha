package route

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/controller"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
	"github.com/alexfalkowski/sasha/internal/site/articles/view"
)

// Register books.
func Register(repo repository.Repository) error {
	errorView, errorPartialView := view.NewErrorView()
	articlesView, articlesPartialView := view.NewArticlesView()
	articleView, articlePartialView := view.NewArticleView()

	mvc.Get("/articles", controller.NewArticlesController(repo, articlesView, errorView))
	mvc.Put("/articles", controller.NewArticlesController(repo, articlesPartialView, errorPartialView))
	mvc.Get("/article/{slug}", controller.NewArticleController(repo, articleView, errorView))
	mvc.Put("/article/{slug}", controller.NewArticleController(repo, articlePartialView, errorPartialView))

	return nil
}
