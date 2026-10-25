package main

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"uuid"

	"github.com/mifumofu/nyow"
	"github.com/urfave/cli/v3"
)

var ErrNoChanges = errors.New("No changes to entry.")

func init() {
	cmds = append(cmds, &cli.Command{
		Name:      "edit",
		Usage:     "Edit the data of a session.",
		ArgsUsage: "[CATEGORY] [DESCRIPTION] [BEGIN] [END]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "id",
				UsageText: "ID of the session, shown in the list subcommand.",
				Required:  true,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "category",
				Usage: "Change the category of the session",
			},
			&cli.StringFlag{
				Name:  "description",
				Usage: "Change the description of the session",
			},
			&cli.TimestampFlag{
				Name:  "start",
				Usage: "Change the start time of the session. Format: " + inputDateFormat,
				Config: cli.TimestampConfig{
					Layouts: []string{inputDateFormat},
				},
			},
			&cli.TimestampFlag{
				Name:  "end",
				Usage: "Change the end time of the session. Format: " + inputDateFormat,
				Config: cli.TimestampConfig{
					Layouts: []string{inputDateFormat},
				},
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

			// update properties within flag
			flags := c.FlagNames()
			if len(flags) == 0 {
				return ErrNoChanges
			}
			orig := s

			if slices.Contains(flags, "category") {
				s.Category = nyow.Category(c.String("category"))
			}

			if slices.Contains(flags, "description") {
				s.Description = c.String("description")
			}

			if slices.Contains(flags, "start") {
				s.Start = c.Timestamp("start")
			}

			if slices.Contains(flags, "end") {
				s.End = c.Timestamp("end")
			}

			// generate summary of edit
			var b strings.Builder
			b.WriteString(lineDiff("Category", string(orig.Category), string(s.Category)))
			b.WriteString(lineDiff("Description", orig.Description, s.Description))
			b.WriteString(lineDiff("Start", orig.Start.Format(displayDateFormat), s.Start.Format(displayDateFormat)))
			b.WriteString(lineDiff("End", orig.End.Format(displayDateFormat), s.End.Format(displayDateFormat)))

			err = provider.Edit(u, s)
			if err != nil {
				return err
			}

			fmt.Println("Updated the following properties:", b.String())
			return nil
		},
	})
}
