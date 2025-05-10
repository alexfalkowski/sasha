package model

import (
	"html/template"

	"github.com/alexfalkowski/sasha/internal/site/meta"
)

// Article for site.
type Article struct {
	*meta.Info `yaml:"-"`
	Meta       meta.Map      `yaml:"-"`
	Body       template.HTML `yaml:"-"`
	Name       string        `yaml:"name,omitempty"`
	Slug       string        `yaml:"slug,omitempty"`
}
