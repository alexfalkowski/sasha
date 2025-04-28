package site

import (
	"embed"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed **/*.tmpl
//go:embed **/*.yaml
//go:embed **/*.txt
//go:embed **/*.jpg
var filesystem embed.FS

// NewFileSystem for site.
func NewFileSystem() fs.FS {
	return filesystem
}

// NewLayout for site.
func NewLayout() mvc.Layout {
	return mvc.Layout("root/layout.tmpl")
}
