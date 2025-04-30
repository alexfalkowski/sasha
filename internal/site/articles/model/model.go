package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

type (
	// Articles for site.
	Articles struct {
		*meta.Info `yaml:"-"`
		Articles   []*Article `yaml:"articles,omitempty"`
	}

	// Article for site.
	Article struct {
		*meta.Info `yaml:"-"`
		Name       string   `yaml:"name,omitempty"`
		Body       string   `yaml:"body,omitempty"`
		Slug       string   `yaml:"slug,omitempty"`
		Images     []*Image `yaml:"images,omitempty"`
	}

	// Image for site.
	Image struct {
		Name        string `yaml:"name,omitempty"`
		Description string `yaml:"description,omitempty"`
	}

	// Error for site.
	Error struct {
		*meta.Info
		Err error
	}
)

// Error satisfies the error interface.
func (e *Error) Error() string {
	return e.Err.Error()
}
