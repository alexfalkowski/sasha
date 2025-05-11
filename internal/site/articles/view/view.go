package view

import (
	"fmt"
	"html/template"
	"io"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/articles/config"
	"github.com/go-sprout/sprout"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

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

// NewRegistries for view.
func NewRegistries(registry *ArticleRegistry) []sprout.Registry {
	return []sprout.Registry{registry}
}

// NewArticleRegistry for view.
func NewArticleRegistry(config *config.Config) *ArticleRegistry {
	return &ArticleRegistry{config: config}
}

// ArticleRegistry for view.
type ArticleRegistry struct {
	handler sprout.Handler
	config  *config.Config
}

// UID provides a unique identifier for the registry.
func (r *ArticleRegistry) UID() string {
	return "alexfalkowski/sasha.article"
}

// LinkHandler connects the Handler to the registry, enabling runtime functionalities.
func (r *ArticleRegistry) LinkHandler(fh sprout.Handler) error {
	r.handler = fh

	return nil
}

// RegisterFunctions adds the provided functions into the given function map.
// This method is called by the Handler to register all functions of a registry.
func (r *ArticleRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "renderBody", r.RenderBody)

	return nil
}

// RenderBody of the article.
// #nosec G203
func (r *ArticleRegistry) RenderBody(slug string, body []byte) template.HTML {
	extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)
	document := parser.Parse(body)

	flags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags:          flags,
		RenderNodeHook: r.image(slug),
	}
	renderer := html.NewRenderer(opts)

	return template.HTML(string(markdown.Render(document, renderer)))
}

func (r *ArticleRegistry) image(slug string) html.RenderNodeFunc {
	return func(writer io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		image, ok := node.(*ast.Image)
		if !ok {
			return ast.GoToNext, false
		}

		if entering {
			src := fmt.Sprintf("%s/%s/%s", r.config.Address, slug, string(image.Destination))
			alt := string(image.Title)
			_, _ = fmt.Fprintf(writer, `<img src="%s" alt="%s">`, src, alt)
		} else {
			_, _ = fmt.Fprint(writer, "</img>")
		}

		return ast.GoToNext, true
	}
}
