package cli

import (
	"github.com/secrethub/secrethub-cli/internals/cli"
	"github.com/secrethub/secrethub-cli/internals/cli/ui"
)

func Run() error {
	io := ui.NewUserIO()
	app := cli.NewApp("demo", "")
	NewServeCommand(io).Register(app)

	err := app.Root.Cmd.Execute()
	return err
}
