package cmd

import "github.com/alexfalkowski/go-service/v2/cli"

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	flags := command.AddServer("server", "Start sasha server", Module)

	flags.AddInput("")
}
