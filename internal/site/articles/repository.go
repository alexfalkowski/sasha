package articles

import (
	"bytes"
	"cmp"
	"io/fs"
	"slices"

	"github.com/alexfalkowski/go-service/encoding/yaml"
	"github.com/alexfalkowski/go-service/runtime"
)

// Repository for books.
type Repository interface {
	// GetArticles from storage.
	GetArticles() *Model

	// GetArticle by slug.
	GetArticle(slug string) *Article
}

// NewRepository for books.
func NewRepository(filesystem fs.FS, enc *yaml.Encoder) Repository {
	return &FSRepository{filesystem: filesystem, enc: enc}
}

// FSRepository has books in a file.
type FSRepository struct {
	filesystem fs.FS
	enc        *yaml.Encoder
}

// GetArticles from a file.
func (r *FSRepository) GetArticles() *Model {
	articles, err := fs.ReadFile(r.filesystem, "articles/articles.yaml")
	runtime.Must(err)

	var m Model
	ptr := &m

	err = r.enc.Decode(bytes.NewBuffer(articles), ptr)
	runtime.Must(err)

	slices.SortFunc(ptr.Articles, func(a, b *Article) int {
		return cmp.Compare(a.Title, b.Title)
	})

	return ptr
}

// GetArticle by slug.
func (r *FSRepository) GetArticle(slug string) *Article {
	articles := r.GetArticles().Articles

	index := slices.IndexFunc(articles, func(a *Article) bool { return a.Slug == slug })
	if index == -1 {
		return nil
	}

	return articles[index]
}
