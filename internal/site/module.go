package site

import (
	"github.com/alexfalkowski/sasha/internal/site/robots"
	"github.com/alexfalkowski/sasha/internal/site/root"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	robots.Module,
	root.Module,
	fx.Provide(NewFS),
	fx.Provide(NewPatterns),
)
