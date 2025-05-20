package view

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// NewRoot view.
func NewRoot() (*mvc.View, *mvc.View) {
	return mvc.NewViewPair("root/view/root.tmpl")
}
