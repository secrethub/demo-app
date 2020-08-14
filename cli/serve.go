package cli

import (
	"fmt"

	"github.com/secrethub/demo-app/app"

	"github.com/secrethub/secrethub-cli/internals/cli/ui"
	"github.com/secrethub/secrethub-cli/internals/secrethub/command"
)

type ServeCommand struct {
	io ui.IO

	host string
	port int

	altPage string
}

func NewServeCommand(io ui.IO) *ServeCommand {
	return &ServeCommand{
		io: io,
	}
}

// Register registers the command, arguments and flags on the provided Registerer.
func (cmd *ServeCommand) Register(r command.Registerer) {
	clause := r.CreateCommand("serve", "Runs the secrethub example by serving a web page.")

	clause.StringVarP(&cmd.host, "host", "h", "127.0.0.1", "The host to serve the webpage on", false, false)
	clause.IntVar(&cmd.port, "port", 8080, "The port to serve the webpage on", false, false)
	clause.StringVar(&cmd.altPage, "alt-page", "", "Path to alternative page file to serve", false, false)

	command.BindAction(clause, nil, cmd.Run)
}

// Run handles the command with the options as specified in the command.
func (cmd *ServeCommand) Run() error {
	fmt.Fprintf(cmd.io.Stdout(), "Serving example app on http://%s:%d\n", cmd.host, cmd.port)
	return app.NewServer(cmd.host, cmd.port, cmd.altPage).Serve()
}
