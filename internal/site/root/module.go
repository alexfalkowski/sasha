package root

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/sasha/internal/site/root/route"
)

// Module for fx.
var Module = di.Module(
	route.Module,
)
