package site

import articles "github.com/alexfalkowski/sasha/internal/site/articles/config"

// Config for the site.
type Config struct {
	Articles *articles.Config `yaml:"articles,omitempty" json:"articles,omitempty" toml:"articles,omitempty"`
}
