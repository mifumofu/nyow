package nyow

import (
	"errors"
	"time"
	"uuid"
)

type Category string

var ErrNotExist = errors.New("This entry does not exist")

type Item struct {
	ID uuid.UUID
	Category
	Description string
	Start       time.Time
	End         time.Time
}

type Provider interface {
	Add(Item) error
	Remove(uuid.UUID) error
	List(after time.Time) []Item
}
