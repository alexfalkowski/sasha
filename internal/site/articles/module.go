package articles

import (
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
	"github.com/alexfalkowski/sasha/internal/site/articles/rest"
	"github.com/alexfalkowski/sasha/internal/site/articles/route"
	"github.com/alexfalkowski/sasha/internal/site/articles/view"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	view.Module,
	rest.Module,
	repository.Module,
	route.Module,
)
