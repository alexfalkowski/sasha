package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/sasha/internal/config"
	"github.com/alexfalkowski/sasha/internal/health"
)

// RegisterServer for cmd.
func RegisterServer(command *cmd.Command) {
	flags := command.AddServer("server", "Start sasha server",
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module,
		config.Module, health.Module, cmd.Module,
	)
	flags.AddInput("")
}
