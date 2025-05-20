package config

import "github.com/alexfalkowski/go-service/v2/client"

// Config for articles.
type Config struct {
	*client.Config `yaml:",inline" json:",inline" toml:",inline"`
}
