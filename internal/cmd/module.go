package cmd

import (
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/sasha/internal/config"
	"github.com/alexfalkowski/sasha/internal/health"
	"github.com/alexfalkowski/sasha/internal/site"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Server,
	config.Module,
	health.Module,
	site.Module,
)
