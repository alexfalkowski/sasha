package articles

import "github.com/alexfalkowski/sasha/internal/site/meta"

type (
	// Model for articles.
	Model struct {
		*meta.Info `yaml:"-"`
		Articles   []*Article `yaml:"articles,omitempty"`
	}

	// Article for our site.
	Article struct {
		*meta.Info `yaml:"-"`
		Name       string   `yaml:"name,omitempty"`
		Body       string   `yaml:"body,omitempty"`
		Slug       string   `yaml:"slug,omitempty"`
		Images     []*Image `yaml:"images,omitempty"`
	}

	// Image for article.
	Image struct {
		Name        string `yaml:"name,omitempty"`
		Description string `yaml:"description,omitempty"`
	}
)
