package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cache"
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/debug"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/go-service/v2/telemetry"
	"github.com/alexfalkowski/go-service/v2/transport"
	"github.com/alexfalkowski/sasha/internal/config"
	"github.com/alexfalkowski/sasha/internal/health"
	"github.com/alexfalkowski/sasha/internal/site"
)

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	flags := command.AddServer("server", "Start sasha server",
		module.Module, debug.Module, telemetry.Module,
		transport.Module, config.Module, health.Module,
		cache.Module, cli.Module, site.Module,
	)
	flags.AddInput("")
}
