package route

import (
	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/controller"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
	"github.com/alexfalkowski/sasha/internal/site/articles/view"
)

// Register books.
func Register(repo repository.Repository) error {
	errorView, errorPartialView := view.NewError()
	articlesView, articlesPartialView := view.NewArticles()
	articleView, articlePartialView := view.NewArticle()

	mvc.Get("/articles", controller.NewArticles(repo, articlesView, errorView))
	mvc.Put("/articles", controller.NewArticles(repo, articlesPartialView, errorPartialView))
	mvc.Get("/article/{slug}", controller.NewArticle(repo, articleView, errorView))
	mvc.Put("/article/{slug}", controller.NewArticle(repo, articlePartialView, errorPartialView))

	return nil
}
