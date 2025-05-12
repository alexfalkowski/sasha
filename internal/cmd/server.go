package cmd

import (
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/sasha/internal/config"
	"github.com/alexfalkowski/sasha/internal/health"
	"github.com/alexfalkowski/sasha/internal/site"
)

// RegisterServer for cmd.
func RegisterServer(command *cmd.Command) {
	flags := command.AddServer("server", "Start sasha server",
		module.Module, debug.Module, telemetry.Module,
		transport.Module, config.Module, health.Module,
		cache.Module, cmd.Module, site.Module,
	)
	flags.AddInput("")
}
