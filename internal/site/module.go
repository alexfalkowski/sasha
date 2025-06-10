package site

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/sasha/internal/site/articles"
	"github.com/alexfalkowski/sasha/internal/site/meta"
	"github.com/alexfalkowski/sasha/internal/site/robots"
	"github.com/alexfalkowski/sasha/internal/site/root"
)

// Module for fx.
var Module = di.Module(
	meta.Module,
	robots.Module,
	root.Module,
	articles.Module,
	di.Constructor(NewFileSystem),
	di.Constructor(NewLayout),
)
