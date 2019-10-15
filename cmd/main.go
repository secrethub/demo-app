package main

import (
	"fmt"
	"os"

	"github.com/secrethub/demo-app/cli"
)

func main() {
	err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "encountered an error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
