package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

// Article for site.
type Article struct {
	*meta.Info `yaml:"-"`
	Meta       meta.Map `yaml:"-"`
	Name       string   `yaml:"name,omitempty"`
	Body       string   `yaml:"body,omitempty"`
	Slug       string   `yaml:"slug,omitempty"`
	Images     []*Image `yaml:"images,omitempty"`
}
