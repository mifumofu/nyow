package main

import (
	"context"
	"fmt"
	"uuid"

	"github.com/mifumofu/nyow/internal/storage"
	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "remove",
		Aliases:   []string{"delete", "r"},
		Usage:     "Remove a session",
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

			fmt.Printf("removed session %v\n", u)
			return provider.Remove(u)
		},
		After: func(ctx context.Context, c *cli.Command) error {
			p, ok := provider.(*storage.JSON)
			if ok {
				return storage.SaveState(*p, "nyow")
			}

			return nil
		},
	})
}
