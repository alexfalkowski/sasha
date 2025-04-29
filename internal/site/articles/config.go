package articles

import "github.com/alexfalkowski/go-service/client"

// Config for articles.
type Config struct {
	*client.Config `yaml:",inline" json:",inline" toml:",inline"`
}
