package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"runtime"
	"time"
	"uuid"

	"github.com/mifumofu/nyow"
)

// save the user's state to a location based on their platform's standards
func SaveState(i JSON, appname string) error {
	sd, err := userStateFile(appname)
	if err != nil {
		return err
	}

	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(sd), 0777)
	if err != nil {
		return err
	}

	return os.WriteFile(sd, b, 0666)
}

// loads the user's state from a file based on their platform's standards
func LoadState(appname string) (nyow.Provider, error) {
	sd, err := userStateFile(appname)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(sd)
	if errors.Is(err, os.ErrNotExist) {
		return &JSON{
			Items:    make(map[uuid.UUID]nyow.Session),
			Trackers: make(map[nyow.Category]time.Time),
		}, nil

	} else if err != nil {
		return nil, err
	}

	var d JSON
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}

	// make if nil after unmarshalling so it doesn't get
	// overriden with nil
	if d.Items == nil {
		d.Items = make(map[uuid.UUID]nyow.Session)
	}
	if d.Trackers == nil {
		d.Trackers = make(map[nyow.Category]time.Time)
	}

	return &d, err
}

// Gets the standard state directory for the user's platform.
// Respects XDG Base Directory Specification on Linux
func userStateDir(appname string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}

	switch runtime.GOOS {
	case "linux":
		// respect XDG base directory specification
		if sh := os.Getenv("XDG_STATE_HOME"); sh != "" {
			return path.Join(sh, appname), nil
		}
		return path.Join(home, ".local", "state", appname), nil

	case "darwin":
		return path.Join(home, "Library", "Application Support", appname), nil
	}

	panic(errors.New("Cannot find state directory for platform!"))
}

// Gets the file used for storing state in
func userStateFile(appname string) (string, error) {
	sd, err := userStateDir(appname)
	if err != nil {
		return "", err
	}

	return path.Join(sd, "state.json"), nil
}
