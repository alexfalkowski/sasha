package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/sasha/internal/config"
	"github.com/alexfalkowski/sasha/internal/health"
	"github.com/alexfalkowski/sasha/internal/site"
)

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	flags := command.AddServer("server", "Start sasha server",
		module.Server,
		config.Module,
		health.Module,
		site.Module,
	)
	flags.AddInput("")
}
