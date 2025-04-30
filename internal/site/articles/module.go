package articles

import (
	"github.com/alexfalkowski/sasha/internal/site/articles/client"
	"github.com/alexfalkowski/sasha/internal/site/articles/repository"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	client.Module,
	repository.Module,
	fx.Invoke(Register),
)
