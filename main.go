package nyow

import (
	"errors"
	"time"
	"uuid"
)

type Category string

var ErrNotExist = errors.New("This entry does not exist")
var ErrTrackAlreadyStarted = errors.New("This current tracker has already been started.")

type Session struct {
	ID uuid.UUID
	Category
	Description string
	Start       time.Time
	End         time.Time
}

func (i Session) Duration() time.Duration {
	return i.End.Sub(i.Start)
}

type Goal struct {
	ID uuid.UUID
	Category
	Amount time.Duration
	Begin  time.Time
}

type GoalProgress struct {
	Goal
	Progress time.Duration
}

type Provider interface {
	Add(Session) error
	Remove(uuid.UUID) error
	Edit(uuid.UUID, Session) error
	Get(uuid.UUID) (Session, error)
	List(*ListOptions) []Session

	// starts a tracker on a category
	TrackStart(Category) error

	// ends a tracker for a category and provide description for session
	TrackEnd(Category, string) (Session, error)

	// cancels a tracking session without saving it
	TrackCancel(Category) error

	// lists all ongoing tracking sessions
	TrackList() (map[Category]time.Time, error)
}

type ListOptions struct {
	Before   *time.Time
	After    *time.Time
	Category *[]Category
}
