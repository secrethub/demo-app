package cli

import (
	"github.com/secrethub/secrethub-cli/internals/cli"
	"github.com/secrethub/secrethub-cli/internals/cli/ui"
)

func Run(args []string) error {
	io := ui.NewUserIO()
	app := cli.NewApp("demo", "")
	NewServeCommand(io).Register(app)

	// Parse also executes the command when parsing is successful.
	_, err := app.Parse(args)
	return err
}
