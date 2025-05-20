package controller

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/net/http/mvc"
	"github.com/alexfalkowski/sasha/internal/site/meta"
	"github.com/alexfalkowski/sasha/internal/site/root/model"
)

// NewRoot controller.
func NewRoot(info *meta.Info, rootView *mvc.View) mvc.Controller[model.Root] {
	return func(_ context.Context) (*mvc.View, *model.Root, error) {
		root := &model.Root{Info: info}

		return rootView, root, nil
	}
}
