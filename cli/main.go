package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mifumofu/nyow"
	"github.com/mifumofu/nyow/internal/storage"
	"github.com/urfave/cli/v3"
)

var cmds []*cli.Command
var provider nyow.Provider

func main() {
	cmd := &cli.Command{
		Name:     "nyow",
		Usage:    "A dead-simple time tracking CLI.",
		Commands: cmds,
	}

	var err error
	provider, err = storage.NewFile(cmd.Name)
	if err != nil {
		fmt.Println(formatError(err))
		os.Exit(1)
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(formatError(err))
		os.Exit(1)
	}
}
