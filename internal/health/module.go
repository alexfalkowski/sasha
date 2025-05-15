package health

import (
	"github.com/alexfalkowski/go-service/health"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewHealthObserver),
	fx.Provide(NewLivenessObserver),
	fx.Provide(NewReadinessObserver),
	fx.Provide(NewRegistrations),
	health.Module,
)
