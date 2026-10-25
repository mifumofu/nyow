package main

import (
	"context"
	"fmt"
	"uuid"

	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "remove",
		Aliases:   []string{"delete", "r"},
		Usage:     "Delete a session",
		ArgsUsage: "[ID]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:     "id",
				Required: true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			u, err := uuid.Parse(c.StringArg("id"))
			if err != nil {
				return err
			}

			fmt.Printf("Removed session %v%v%v.\n", colorInfo, u, colorReset)
			return provider.Remove(u)
		},
	})
}
