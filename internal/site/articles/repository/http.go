package repository

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"

	se "github.com/alexfalkowski/go-service/errors"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/mime"
	"github.com/alexfalkowski/sasha/internal/site/articles/client"
	articles "github.com/alexfalkowski/sasha/internal/site/articles/config"
	"github.com/alexfalkowski/sasha/internal/site/articles/model"
	sm "github.com/alexfalkowski/sasha/internal/site/meta"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// NewRepository for books.
func NewRepository(info *sm.Info, config *articles.Config, client *client.Client) Repository {
	return &HTTPRepository{info: info, config: config, client: client}
}

// HTTPRepository uses a client to get from a site (public bucket).
type HTTPRepository struct {
	info   *sm.Info
	config *articles.Config
	client *client.Client
}

// GetArticles from the public bucket.
func (r *HTTPRepository) GetArticles(ctx context.Context) (*model.Articles, error) {
	articles := &model.Articles{}
	opts := &client.Options{
		ContentType: mime.YAMLMediaType,
		Response:    articles,
	}
	url := r.config.Address + "/articles.yml"

	if err := r.client.Get(ctx, url, opts); err != nil {
		if client.IsNotFound(err) {
			articles.Meta = meta.Strings(ctx, "")
			articles.Info = r.info

			return articles, nil
		}

		err := &model.Error{
			Meta: meta.Strings(ctx, ""),
			Info: r.info,
			Err:  se.Prefix("repository: get articles", err),
		}

		return nil, err
	}

	articles.Meta = meta.Strings(ctx, "")
	articles.Info = r.info

	return articles, nil
}

// GetArticle by slug from the public bucket.
// #nosec G203
func (r *HTTPRepository) GetArticle(ctx context.Context, slug string) (*model.Article, error) {
	article, err := r.getArticleConfig(ctx, slug)
	if err != nil {
		return nil, err
	}

	body, err := r.getArticleBody(ctx, slug)
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)
	document := parser.Parse(body)

	flags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Title:          article.Name,
		Flags:          flags,
		RenderNodeHook: r.image(slug),
	}
	renderer := html.NewRenderer(opts)

	article.Meta = meta.Strings(ctx, "")
	article.Info = r.info
	article.Body = template.HTML(string(markdown.Render(document, renderer)))

	return article, nil
}

func (r *HTTPRepository) getArticleConfig(ctx context.Context, slug string) (*model.Article, error) {
	article := &model.Article{}
	opts := &client.Options{
		ContentType: mime.YAMLMediaType,
		Response:    article,
	}
	url := fmt.Sprintf("%s/%s/article.yml", r.config.Address, slug)

	if err := r.get(ctx, url, opts); err != nil {
		return nil, err
	}

	return article, nil
}

func (r *HTTPRepository) getArticleBody(ctx context.Context, slug string) ([]byte, error) {
	buffer := &bytes.Buffer{}
	opts := &client.Options{
		ContentType: mime.MarkdownMediaType,
		Response:    buffer,
	}
	url := fmt.Sprintf("%s/%s/article.md", r.config.Address, slug)

	if err := r.get(ctx, url, opts); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (r *HTTPRepository) get(ctx context.Context, url string, opts *client.Options) error {
	if err := r.client.Get(ctx, url, opts); err != nil {
		if client.IsNotFound(err) {
			err = ErrNotFound
		}

		err := &model.Error{
			Meta: meta.Strings(ctx, ""),
			Info: r.info,
			Err:  se.Prefix("repository: get url", err),
		}

		return err
	}

	return nil
}

func (r *HTTPRepository) image(slug string) html.RenderNodeFunc {
	return func(writer io.Writer, node ast.Node, _ bool) (ast.WalkStatus, bool) {
		image, ok := node.(*ast.Image)
		if !ok {
			return ast.GoToNext, false
		}

		src := fmt.Sprintf("%s/%s/%s", r.config.Address, slug, string(image.Destination))
		alt := string(image.Title)
		_, _ = fmt.Fprintf(writer, `<img src="%s" alt="%s" />`, src, alt)

		return ast.GoToNext, true
	}
}
