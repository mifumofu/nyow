package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

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

func init() {
	cmds = append(cmds, &cli.Command{
		Name:  "version",
		Usage: "Versioning information",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Output JSON.",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			d, ok := debug.ReadBuildInfo()
			if !ok {
				return errors.New("Failed to read build information from the binary")
			}

			var b strings.Builder
			fmt.Fprintf(&b, "%v (%v)", d.Path, d.Main.Version)
			fmt.Fprintf(&b, "\ncompiled with %v", d.GoVersion)

			b.WriteString("\nbuild flags: ")
			for _, v := range d.Settings {
				fmt.Fprintf(&b, "\n\t%v = %v", v.Key, v.Value)
			}

			fmt.Println(b.String())

			return nil
		},
	})
}
