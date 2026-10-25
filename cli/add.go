package main

import (
	"context"
	"fmt"

	"github.com/mifumofu/nyow"
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
				UsageText: "Category of the session",
				Required:  true,
			},
			&cli.StringArg{
				Name:      "description",
				UsageText: "Description of the session",
				Required:  true,
			},
			&cli.TimestampArg{
				Name:      "begin",
				UsageText: "Beginning time of the session",
				Required:  true,
				Config: cli.TimestampConfig{
					Layouts: []string{inputDateFormat},
				},
			},
			&cli.TimestampArg{
				Name:      "end",
				UsageText: "End time of the session",
				Required:  true,
				Config: cli.TimestampConfig{
					Layouts: []string{inputDateFormat},
				},
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			item := nyow.Session{
				Category:    nyow.Category(c.StringArg("category")),
				Description: c.StringArg("description"),
				Start:       c.TimestampArg("begin"),
				End:         c.TimestampArg("end"),
			}

			fmt.Printf("Created session with description %v%v%v of category %v%v%v.\n", colorInfo, item.Description, colorReset, colorInfo, item.Category, colorReset)
			return provider.Add(item)
		},
	})
}
