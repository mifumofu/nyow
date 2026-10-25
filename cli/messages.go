package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mifumofu/nyow"
)

var displayDateFormat = "Monday, 02 Jan 2006 3:04PM"
var inputDateFormat = "2006-01-02 15:04"

var colorPrimary = "\x1b[45m\x1b[30m"
var colorDim = "\x1b[2m"
var colorBold = "\x1b[1m"
var colorReset = "\x1b[0m"

var colorRed = "\x1b[31m"
var colorGreen = "\x1b[32m"
var colorYellow = "\x1b[33m"

func printItems(i []nyow.Session) string {
	var b strings.Builder
	for _, x := range i {
		b.WriteString(printItem(x))
		b.WriteRune('\n')
		b.WriteRune('\n')
	}
	return b.String()
}

// prints out a formatted Item
func printItem(i nyow.Session) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%v%v ", colorBold, i.Description)
	fmt.Fprintf(&b, "%v %v %v", colorPrimary, string(i.Category), colorReset)
	fmt.Fprintf(&b, " %v%v", colorDim, i.ID)

	fmt.Fprintf(&b, "\n%vAt %v until %v for %v.", colorReset, i.Start.Format(displayDateFormat), i.End.Format(displayDateFormat), i.Duration().Truncate(time.Second))
	return b.String()
}

func lineDiff(description, before, after string) string {
	if before != after {
		return fmt.Sprintf("\n%v- %v%v%v: %v%v %v-> %v%v", colorReset, colorBold, description, colorReset, colorRed, before, colorReset, colorGreen, after)
	}

	return fmt.Sprintf("\n%v- %v%v%v: %v%v", colorReset, colorBold, description, colorReset, colorYellow, after)
}

// Formats an error into something more user friendly
func formatError(err error) string {
	return fmt.Sprintf("%v%v%vERROR%v: %v", colorReset, colorRed, colorBold, colorReset, err)
}
