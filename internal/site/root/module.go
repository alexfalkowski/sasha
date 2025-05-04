package root

import (
	"github.com/alexfalkowski/sasha/internal/site/root/route"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	route.Module,
)
