package view

import "github.com/alexfalkowski/go-service/net/http/mvc"

// NewErrorView for articles.
func NewErrorView() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("articles/view/error.tmpl")
}

// NewArticleView for articles.
func NewArticleView() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("articles/view/article.tmpl")
}

// NewArticlesView for articles.
func NewArticlesView() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("articles/view/articles.tmpl")
}
