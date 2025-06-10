package cmd

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/sasha/internal/config"
	"github.com/alexfalkowski/sasha/internal/health"
	"github.com/alexfalkowski/sasha/internal/site"
)

// Module for fx.
var Module = di.Module(
	module.Server,
	config.Module,
	health.Module,
	site.Module,
)
