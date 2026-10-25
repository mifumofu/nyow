package main

import (
	"context"
	"encoding/json"
	"fmt"
	"uuid"

	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "get",
		Aliases:   []string{"show"},
		Usage:     "Get a session",
		ArgsUsage: "<ID>",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "id",
				UsageText: "ID of session",
				Required:  true,
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Output JSON.",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			u, err := uuid.Parse(c.StringArg("id"))
			if err != nil {
				return err
			}

			s, err := provider.Get(u)
			if err != nil {
				return err
			}

			if c.Bool("json") {
				b, err := json.MarshalIndent(s, "", "\t")
				if err != nil {
					return err
				}
				fmt.Println(string(b))
				return nil
			}

			fmt.Println(printItem(s))
			return nil
		},
	})
}
