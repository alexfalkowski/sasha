package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

// Articles for site.
type Articles struct {
	*meta.Info `yaml:"-"`
	Meta       meta.Map   `yaml:"-"`
	Articles   []*Article `yaml:"articles,omitempty"`
}
