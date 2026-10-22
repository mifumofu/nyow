package storage

import (
	"maps"
	"slices"
	"time"
	"uuid"

	"github.com/mifumofu/nyow"
)

type JSON struct {
	items map[uuid.UUID]nyow.Item
}

func (j *JSON) Add(i nyow.Item) error {
	if i.ID == uuid.Nil() {
		i.ID = uuid.New()
	}
	j.items[i.ID] = i
	return nil
}

func (j *JSON) Remove(u uuid.UUID) error {
	delete(j.items, u)
	return nil
}

func (j *JSON) List(after time.Time) []nyow.Item {
	s := slices.Collect(maps.Values(j.items))

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
