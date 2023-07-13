package flag_icons

import (
	"embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/rs/zerolog"
	"io"
	"sync"
)

var (
	mutex  = sync.Mutex{}
	logger = zerolog.Nop()
	cache  = map[string]fyne.Resource{}

	//go:embed flags
	resources embed.FS
)

func RegisterLogger(l zerolog.Logger) {
	logger = l
}

func Icon(name string) (fyne.Resource, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if i, exists := cache[name]; exists {
		return i, nil
	}
	if f, err := resources.Open(fmt.Sprintf("flags/%s.svg", name)); err != nil {
		return nil, err
	} else if icon, readErr := io.ReadAll(f); readErr != nil {
		return nil, readErr
	} else {
		cache[name] = fyne.NewStaticResource(name, icon)
		logger.Debug().Str("name", name).Msg("flag-resource has been loaded + cached")

		return cache[name], nil
	}

}

func MustIcon(name string) fyne.Resource {
	if icon, err := Icon(name); err != nil {
		logger.Warn().Err(err).Str("name", name).Msg("could not find flag-resource")

		return theme.ErrorIcon()
	} else {
		return icon
	}
}
