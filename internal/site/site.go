package site

import (
	"embed"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed **/*.tmpl
//go:embed **/*.yaml
//go:embed **/*.txt
var filesystem embed.FS

// NewFS for site.
func NewFS() fs.FS {
	return filesystem
}

// NewPatterns for site.
func NewPatterns() mvc.Patterns {
	return mvc.Patterns{"**/*.tmpl"}
}
