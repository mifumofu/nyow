package main

import (
	"context"
	"fmt"

	"github.com/mifumofu/nyow"
	"github.com/mifumofu/nyow/internal/storage"
	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "add",
		Aliases:   []string{"a"},
		Usage:     "Manually add a session",
		ArgsUsage: "[CATEGORY] [DESCRIPTION] [BEGIN] [END]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "category",
				UsageText: "Category of the session.",
				Required:  true,
			},
			&cli.StringArg{
				Name:      "description",
				UsageText: "Description of the session",
				Required:  true,
			},
			&cli.StringArg{
				Name:      "begin",
				UsageText: "Beginning time of the session",
				Required:  true,
			},
			&cli.StringArg{
				Name:      "end",
				UsageText: "End time of the session",
				Required:  true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			item := nyow.Session{
				Category:    nyow.Category(c.StringArg("category")),
				Description: c.StringArg("description"),
				Start:       c.TimestampArg("begin"),
				End:         c.TimestampArg("end"),
			}

			fmt.Printf("added session with description \"%v\" of category \"%v\"", item.Description, item.Category)
			return provider.Add(item)
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
