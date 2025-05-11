package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

// Article for site.
type Article struct {
	*meta.Info `yaml:"-"`
	Name       string `yaml:"name,omitempty"`
	Slug       string `yaml:"slug,omitempty"`
	Body       []byte `yaml:"-"`
}
