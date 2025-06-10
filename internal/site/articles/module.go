package articles

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
	"github.com/alexfalkowski/sasha/internal/site/articles/rest"
	"github.com/alexfalkowski/sasha/internal/site/articles/route"
	"github.com/alexfalkowski/sasha/internal/site/articles/view"
)

// Module for fx.
var Module = di.Module(
	view.Module,
	rest.Module,
	repository.Module,
	route.Module,
)
