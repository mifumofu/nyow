package nyow

import (
	"errors"
	"time"
	"uuid"
)

type Category string

var ErrNotExist = errors.New("This entry does not exist")
var ErrTrackAlreadyStarted = errors.New("This current tracker has already been started.")

type Item struct {
	ID uuid.UUID
	Category
	Description string
	Start       time.Time
	End         time.Time
}

func (i Item) Duration() time.Duration {
	return i.End.Sub(i.Start)
}

type Provider interface {
	Add(Item) error
	Remove(uuid.UUID) error
	List(after time.Time) []Item

	// starts a tracker on a category
	TrackStart(Category) error

	// ends a tracker for a category and provide description for session
	TrackEnd(Category, string) (Item, error)
}
