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
		Name:  "overview",
		Usage: "Get an overview of time over a period",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Output JSON.",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			msg := summaryMessage(provider, time.Hour*24*7, time.Hour*24, time.Hour/2)

			fmt.Println(msg)
			return nil
		},
	})
}

// Gives an overview of sessions through displaying a table
// total is the period that the overview contains
// rowperiod is the period that rows are divided into
// boxperiod is the period that individual boxes represent
func summaryMessage(p nyow.Provider, total, rowperiod, boxperiod time.Duration) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Overview for time spent over %v.\n", bigDurationString(total))
	b.WriteString("\n                                1 1 1 1 1 1 1 1 1 1 2 2 2 2 2")
	b.WriteString("\n              1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4\n")

	pbefore := time.Now()
	pafter := time.Now().Add(-boxperiod)

	// write row key
	for range total / rowperiod {
		fmt.Fprintf(&b, "%v%9v %2v ", colorReset, pbefore.Weekday(), pbefore.Day())

		// write individual blocks, one per subblock
		for range rowperiod / boxperiod {
			pbefore = pbefore.Add(-boxperiod)
			pafter = pafter.Add(-boxperiod)

			i := p.List(&nyow.ListOptions{
				Before: &pbefore,
				After:  &pafter,
			})

			if len(i) == 0 {
				b.WriteString(colorDim)
			} else {
				b.WriteString(colorReset)
			}
			b.WriteString("■")
		}

		b.WriteString("\n")
	}

	return b.String()
}

// fixes a t.String() limitation where hours is the largest unit, goes to a year
// only includes one unit for simplicity eg. no 1 week 4 hours, only 1 week
func bigDurationString(d time.Duration) string {
	u := uint64(d)

	units := []struct {
		Name     string
		Duration uint64
	}{
		{
			Name:     "year",
			Duration: uint64(time.Hour) * 24 * 365,
		},
		{
			Name:     "month",
			Duration: uint64(time.Hour) * 24 * 7 * 30,
		},
		{
			Name:     "week",
			Duration: uint64(time.Hour) * 24 * 7,
		},
		{
			Name:     "day",
			Duration: uint64(time.Hour) * 24,
		},
		{
			Name:     "hour",
			Duration: uint64(time.Hour),
		},
		{
			Name:     "minute",
			Duration: uint64(time.Minute),
		},
		{
			Name:     "second",
			Duration: uint64(time.Second),
		},
	}

	for _, i := range units {
		m := u / i.Duration
		if m == 1 {
			return fmt.Sprintf("%v %v", m, i.Name)
		} else if m > 1 {
			return fmt.Sprintf("%v %vs", m, i.Name)
		}
	}

	// we have already dealt with the problem of big
	// durations being in an odd unit eg. 1 week -> 168hours
	// by this point, so we just fallback to the std library implentation
	// if its small enough to slip by.
	return d.String()
}
