package articles

import (
	"bytes"
	"cmp"
	"io/fs"
	"slices"

	"github.com/alexfalkowski/go-service/encoding/yaml"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/types/ptr"
	"github.com/alexfalkowski/sasha/internal/site/meta"
)

// Repository for books.
type Repository interface {
	// GetArticles from storage.
	GetArticles() *Model

	// GetArticle by slug.
	GetArticle(slug string) *Article
}

// NewRepository for books.
func NewRepository(info *meta.Info, filesystem fs.FS, enc *yaml.Encoder) Repository {
	return &FSRepository{info: info, filesystem: filesystem, enc: enc}
}

// FSRepository has books in a file.
type FSRepository struct {
	info       *meta.Info
	filesystem fs.FS
	enc        *yaml.Encoder
}

// GetArticles from a file.
func (r *FSRepository) GetArticles() *Model {
	articles, err := fs.ReadFile(r.filesystem, "articles/articles.yaml")
	runtime.Must(err)

	model := ptr.Zero[Model]()

	err = r.enc.Decode(bytes.NewBuffer(articles), model)
	runtime.Must(err)

	slices.SortFunc(model.Articles, func(a, b *Article) int {
		return cmp.Compare(a.Name, b.Name)
	})

	model.Info = r.info

	return model
}

// GetArticle by slug.
func (r *FSRepository) GetArticle(slug string) *Article {
	articles := r.GetArticles().Articles

	index := slices.IndexFunc(articles, func(a *Article) bool { return a.Slug == slug })
	if index == -1 {
		return nil
	}

	article := articles[index]
	article.Info = r.info

	return article
}
