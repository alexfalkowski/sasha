package site

import (
	"embed"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed **/*.tmpl
//go:embed articles/view/*tmpl
//go:embed **/*.txt
var filesystem embed.FS

// NewFileSystem for site.
func NewFileSystem() fs.FS {
	return filesystem
}

// NewLayout for site.
func NewLayout() *mvc.Layout {
	return mvc.NewLayout("root/full.tmpl", "root/partial.tmpl")
}
