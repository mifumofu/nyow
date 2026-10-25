package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mifumofu/nyow"
	"github.com/mifumofu/nyow/internal/storage"
	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "start",
		Usage:     "Start tracking a session",
		ArgsUsage: "[CATEGORY]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "category",
				UsageText: "Category of the session.",
				Required:  true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cat := nyow.Category(c.StringArg("category"))
			err := provider.TrackStart(cat)
			if err != nil {
				return err
			}

			fmt.Printf("started tracking session for \"%v\"\n", cat)
			return nil
		},
		After: func(ctx context.Context, c *cli.Command) error {
			p, ok := provider.(*storage.JSON)
			if ok {
				return storage.SaveState(*p, "nyow")
			}
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
				UsageText: "Category of the session.",
				Required:  true,
			},
			&cli.StringArg{
				Name:      "description",
				UsageText: "Description of the session.",
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

			fmt.Printf("ended session for \"%v\" after %v.\n", cat, i.Duration().Round(time.Second))
			return nil
		},
		After: func(ctx context.Context, c *cli.Command) error {
			p, ok := provider.(*storage.JSON)
			if ok {
				return storage.SaveState(*p, "nyow")
			}
			return nil
		},
	}, &cli.Command{
		Name:      "cancel",
		Usage:     "Cancel an ongoing tracking session without saving anything.",
		ArgsUsage: "[CATEGORY]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "category",
				UsageText: "Category of the session.",
				Required:  true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cat := nyow.Category(c.StringArg("category"))

			err := provider.TrackCancel(cat)
			if err != nil {
				return err
			}

			fmt.Printf("Cancelled tracking session of category \"%v\".\n", cat)
			return nil
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
