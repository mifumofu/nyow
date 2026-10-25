package storage

import (
	"maps"
	"slices"
	"time"
	"uuid"

	"github.com/mifumofu/nyow"
)

type JSON struct {
	Items    map[uuid.UUID]nyow.Session
	Trackers map[nyow.Category]time.Time
}

func (j *JSON) Add(i nyow.Session) error {
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

func (j *JSON) List(options *nyow.ListOptions) []nyow.Session {
	s := slices.Collect(maps.Values(j.Items))

	if options == nil {
		return s
	}

	var v []nyow.Session
	for _, i := range s {
		// filter based on occur before
		if options.Before != nil && !i.Start.Before(*options.Before) {
			continue
		}

		// filter based on occur after
		if options.After != nil && !i.Start.After(*options.After) {
			continue
		}

		// filter based on category
		if options.Category != nil && !slices.Contains(*options.Category, i.Category) {
			continue
		}

		v = append(v, i)
	}

	return v
}

func (j *JSON) Get(u uuid.UUID) (nyow.Session, error) {
	if s, ok := j.Items[u]; ok {
		return s, nil
	}
	return nyow.Session{}, nyow.ErrNotExist
}

func (j *JSON) Edit(u uuid.UUID, s nyow.Session) error {
	_, ok := j.Items[u]
	if !ok {
		return nyow.ErrNotExist
	}

	j.Items[u] = s
	return nil
}

func (j *JSON) TrackStart(c nyow.Category) error {
	if _, exists := j.Trackers[c]; exists {
		return nyow.ErrTrackAlreadyStarted
	}

	j.Trackers[c] = time.Now()
	return nil
}

func (j *JSON) TrackEnd(c nyow.Category, desc string) (nyow.Session, error) {
	t, exists := j.Trackers[c]

	if !exists {
		return nyow.Session{}, nyow.ErrNotExist
	}

	i := nyow.Session{
		Start:       t,
		End:         time.Now(),
		Description: desc,
		Category:    c,
	}
	delete(j.Trackers, c)

	return i, j.Add(i)
}

func (j *JSON) TrackCancel(c nyow.Category) error {
	if _, exists := j.Trackers[c]; !exists {
		return nyow.ErrNotExist
	}

	delete(j.Trackers, c)
	return nil
}
