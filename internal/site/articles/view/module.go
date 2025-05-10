package view

import "go.uber.org/fx"

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewArticleRegistry),
	fx.Provide(NewRegistries),
)
