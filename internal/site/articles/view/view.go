package view

import "github.com/alexfalkowski/go-service/net/http/mvc"

// NewError view.
func NewError() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("articles/view/error.tmpl")
}

// NewArticle view.
func NewArticle() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("articles/view/article.tmpl")
}

// NewArticles view.
func NewArticles() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("articles/view/articles.tmpl")
}
