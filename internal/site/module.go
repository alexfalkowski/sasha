package site

import (
	"github.com/alexfalkowski/sasha/internal/site/articles"
	"github.com/alexfalkowski/sasha/internal/site/meta"
	"github.com/alexfalkowski/sasha/internal/site/robots"
	"github.com/alexfalkowski/sasha/internal/site/root"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	meta.Module,
	robots.Module,
	root.Module,
	articles.Module,
	fx.Provide(NewFileSystem),
	fx.Provide(NewLayout),
)
