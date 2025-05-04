package route

import (
	"github.com/alexfalkowski/go-service/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/meta"
	"github.com/alexfalkowski/sasha/internal/site/root/controller"
	"github.com/alexfalkowski/sasha/internal/site/root/view"
)

// Register root.
func Register(info *meta.Info) {
	rootView, rootPartialView := view.NewRoot()

	mvc.Get("/", controller.NewRoot(info, rootView))
	mvc.Put("/", controller.NewRoot(info, rootPartialView))
}
