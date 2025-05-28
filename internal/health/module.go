package health

import "go.uber.org/fx"

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewHealthObserver),
	fx.Provide(NewLivenessObserver),
	fx.Provide(NewReadinessObserver),
	fx.Provide(NewRegistrations),
)
