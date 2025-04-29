package root

import (
	"context"

	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/meta"
)

// Root for root.
type Root struct {
	*meta.Info
}

// Register root.
func Register(info *meta.Info) {
	rootView, rootPartialView := mvc.NewViewPair("root/root.tmpl")

	mvc.Get("/", func(_ context.Context) (*mvc.View, *Root, error) {
		root := &Root{Info: info}

		return rootView, root, nil
	})

	mvc.Put("/", func(_ context.Context) (*mvc.View, *Root, error) {
		root := &Root{Info: info}

		return rootPartialView, root, nil
	})
}
