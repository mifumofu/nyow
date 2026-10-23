package storage

import (
	"maps"
	"slices"
	"time"
	"uuid"

	"github.com/mifumofu/nyow"
)

type JSON struct {
	Items    map[uuid.UUID]nyow.Item
	Trackers map[nyow.Category]time.Time
}

func (j *JSON) Add(i nyow.Item) error {
	if i.ID == uuid.Nil() {
		i.ID = uuid.New()
	}
	j.Items[i.ID] = i
	return nil
}

func (j *JSON) Remove(u uuid.UUID) error {
	delete(j.Items, u)
	return nil
}

func (j *JSON) List(after time.Time) []nyow.Item {
	s := slices.Collect(maps.Values(j.Items))

	// return all entries if after filter not filled in
	if after.IsZero() {
		return s
	}

	var v []nyow.Item
	for _, i := range s {
		if i.Start.After(after) {
			v = append(v, i)
		}
	}

	return v
}

func (j *JSON) TrackStart(c nyow.Category) error {
	if _, exists := j.Trackers[c]; exists {
		return nyow.ErrTrackAlreadyStarted
	}

	j.Trackers[c] = time.Now()
	return nil
}

func (j *JSON) TrackEnd(c nyow.Category, desc string) (nyow.Item, error) {
	t, exists := j.Trackers[c]

	if !exists {
		return nyow.Item{}, nyow.ErrNotExist
	}

	i := nyow.Item{
		Start:       t,
		End:         time.Now(),
		Description: desc,
		Category:    c,
	}
	delete(j.Trackers, c)

	return i, j.Add(i)
}
