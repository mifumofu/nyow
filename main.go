package nyow

import (
	"time"
)

type Category string

type Item struct {
	Category
	Description string
	Start       time.Time
	End         time.Time
}
