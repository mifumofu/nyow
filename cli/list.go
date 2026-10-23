package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mifumofu/nyow"
	"github.com/urfave/cli/v3"
)

func init() {
	cmds = append(cmds, &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List all sessions",
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name:  "before",
				Usage: "Sessions that started before time provided. Format: " + time.DateOnly,
				Config: cli.TimestampConfig{
					Layouts: []string{time.DateOnly},
				},
			},
			&cli.TimestampFlag{
				Name:  "after",
				Usage: "Sessions that started after time provided. Format: " + time.DateTime,
				Config: cli.TimestampConfig{
					Layouts: []string{time.DateOnly},
				},
			},
			&cli.StringFlag{
				Name:  "category",
				Usage: "Sessions in a category, filter multiple by seperating them with commas.",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			after := c.Timestamp("after")
			before := c.Timestamp("before")
			cat := []nyow.Category{}
			for _, x := range strings.Split(c.String("category"), " ") { // cannot do this with []nyow.Category()
				if x == "" {
					break
				}
				cat = append(cat, nyow.Category(x))
			}

			conf := &nyow.ListOptions{}
			var confSet bool = false
			if !before.IsZero() {
				conf.Before = &before
				confSet = true
			}
			if !after.IsZero() {
				conf.After = &after
				confSet = true
			}
			if len(cat) != 0 {
				conf.Category = &cat
				confSet = true
			}

			if !confSet {
				conf = nil
			}

			items := provider.List(conf)
			fmt.Println(printItems(items))
			return nil
		},
	})
}
