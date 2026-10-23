package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mifumofu/nyow"
)

var colorPrimary = "\x1b[45m\x1b[30m"
var colorDim = "\x1b[2m"
var colorBold = "\x1b[1m"
var colorReset = "\x1b[0m"

func printItems(i []nyow.Item) string {
	var b strings.Builder
	for _, x := range i {
		b.WriteString(printItem(x))
		b.WriteRune('\n')
		b.WriteRune('\n')
	}
	return b.String()
}

// prints out a formatted Item
func printItem(i nyow.Item) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%v%v ", colorBold, i.Description)
	fmt.Fprintf(&b, "%v %v %v", colorPrimary, string(i.Category), colorReset)
	fmt.Fprintf(&b, " %v%v", colorDim, i.ID)

	fmt.Fprintf(&b, "\n%vAt %v for %v.", colorReset, i.Start.Format(time.DateTime), i.Duration().Truncate(time.Second))
	return b.String()
}
