package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mifumofu/nyow"
	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "start",
		Usage:     "Start tracking a new session",
		ArgsUsage: "[CATEGORY]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "category",
				UsageText: "Category of the session",
				Required:  true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cat := nyow.Category(c.StringArg("category"))
			err := provider.TrackStart(cat)
			if err != nil {
				return err
			}

			fmt.Printf("Started tracking session for category %v%v%v.\n", colorInfo, cat, colorReset)
			return nil
		},
	}, &cli.Command{
		Name:      "end",
		Usage:     "End a session",
		Aliases:   []string{"stop"},
		ArgsUsage: "[CATEGORY] [DESCRIPTION]",
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
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cat := nyow.Category(c.StringArg("category"))
			desc := c.StringArg("description")
			i, err := provider.TrackEnd(cat, desc)
			if err != nil {
				return err
			}

			fmt.Printf("Ended session for %v%v%v after %v%v.\n", colorInfo, cat, colorReset, colorInfo, i.Duration().Round(time.Second))
			return nil
		},
	}, &cli.Command{
		Name:      "cancel",
		Usage:     "Cancel an ongoing tracking session without saving anything.",
		ArgsUsage: "[CATEGORY]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "category",
				UsageText: "Category of the session",
				Required:  true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cat := nyow.Category(c.StringArg("category"))

			err := provider.TrackCancel(cat)
			if err != nil {
				return err
			}

			fmt.Printf("Cancelled tracking session of category %v%v%v.\n", colorInfo, cat, colorReset)
			return nil
		},
	}, &cli.Command{
		Name:  "trackers",
		Usage: "Lists all ongoing trackers.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Output JSON.",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			trackers, err := provider.TrackList()
			if err != nil {
				return err
			}

			fmt.Println(printsTrackers(trackers))
			return nil
		},
	})
}
