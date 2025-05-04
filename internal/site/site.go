package site

import (
	"embed"
	"io/fs"

	"github.com/alexfalkowski/go-service/net/http/mvc"
)

//go:embed root/view/*.tmpl
//go:embed root/layout/*.tmpl
//go:embed articles/view/*.tmpl
//go:embed robots/robots.txt
var filesystem embed.FS

// NewFileSystem for site.
func NewFileSystem() fs.FS {
	return filesystem
}

// NewLayout for site.
func NewLayout() *mvc.Layout {
	return mvc.NewLayout("root/layout/full.tmpl", "root/layout/partial.tmpl")
}
