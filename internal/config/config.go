package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/sasha/internal/health"
	"github.com/alexfalkowski/sasha/internal/site"
	"github.com/alexfalkowski/sasha/internal/site/articles"
)

// Config for the service.
type Config struct {
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	Site           *site.Config   `yaml:"site,omitempty" json:"site,omitempty" toml:"site,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}

func articlesConfig(cfg *Config) *articles.Config {
	return cfg.Site.Articles
}
