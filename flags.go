package flag_icons

import (
	"embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/rs/zerolog"
	"golang.org/x/text/language"
	"io"
	"strings"
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

// IconForLanguage gibt das Flaggen-Icon für das gegebene language.Tag zurück.
// Es versucht, die Region aus dem Tag zu verwenden (z.B. "US" aus "en-US").
// Wenn keine Region vorhanden ist, gibt es einen Fehler zurück oder ein Standard-Icon.
func IconForLanguage(tag language.Tag) (fyne.Resource, error) {
	// Versuche, die Region aus dem Tag zu extrahieren
	region, confidence := tag.Region()

	// Wenn keine Region sicher bestimmt werden kann, könnten wir einen Fehler zurückgeben
	// oder versuchen, die Basissprache zu verwenden (was aber oft nicht eindeutig ist, z.B. "en").
	// Hier geben wir einen Fehler zurück, wenn keine Region vorhanden ist.
	if confidence == language.No {
		// Alternativ: return MustIcon("xx"), nil // für eine generische Flagge
		return nil, fmt.Errorf("kann keine eindeutige Region für Sprache '%s' bestimmen", tag.String())
	}

	// Wandle den Ländercode (Region) in Kleinbuchstaben um, da die Icons so benannt sind
	countryCode := strings.ToLower(region.String())

	// Verwende die bestehende Icon-Funktion
	return Icon(countryCode)
}

// MustIconForLanguage ist wie IconForLanguage, gibt aber theme.ErrorIcon() bei Fehlern zurück.
func MustIconForLanguage(tag language.Tag) fyne.Resource {
	icon, err := IconForLanguage(tag)
	if err != nil {
		logger.Warn().Err(err).Str("language_tag", tag.String()).Msg("could not find flag-resource for language tag")
		return theme.ErrorIcon()
	}
	return icon
}
