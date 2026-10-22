package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List all sessions",
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name:  "after",
				Usage: "Only return sessions which occurred after date provided",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			t := c.Timestamp("after")
			items := provider.List(t)
			for _, x := range items {
				fmt.Printf("%v - [%v] %v\n", x.ID, x.Category, x.Description)
			}
			return nil
		},
	})
}
