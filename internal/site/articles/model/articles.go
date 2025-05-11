package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

// Articles for site.
type Articles struct {
	Info     *meta.Info `yaml:"-"`
	Articles []*Article `yaml:"articles,omitempty"`
}
