// Package fyne_flagicons provides country flag icons for Fyne apps.
//
// Flags are addressed by their lowercase ISO 3166-1 alpha-2 code (plus a few
// subdivisions such as "gb-eng"); the generated Flag* constants list all of
// them. The SVGs are taken from https://github.com/lipis/flag-icons.
package fyne_flagicons

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/timzifer/fyne_iconkit"
	"golang.org/x/text/language"
)

//go:generate go run github.com/timzifer/fyne_iconkit/cmd/iconconst -dir flags -prefix Flag -upper -out flags_gen.go

var (
	//go:embed flags
	flags embed.FS

	set = fyne_iconkit.NewSet(flags, "flags")
)

// Icon returns the flag for the given country code, e.g. "de" or "gb-eng".
// The code is case-insensitive. Unknown codes yield an error wrapping
// fs.ErrNotExist.
func Icon(code string) (fyne.Resource, error) {
	return set.Icon(strings.ToLower(code))
}

// MustIcon is like Icon but returns theme.ErrorIcon() for unknown codes.
func MustIcon(code string) fyne.Resource {
	return set.MustIcon(strings.ToLower(code))
}

// Source returns the SVG content of the flag for the given country code.
func Source(code string) ([]byte, error) {
	return set.Source(strings.ToLower(code))
}

// PNG rasterizes the flag for the given country code. size is the length of
// the longer edge in pixels.
func PNG(code string, size int) ([]byte, error) {
	return set.PNG(strings.ToLower(code), size)
}

// Names returns the codes of all available flags.
func Names() []string {
	return set.Names()
}

// IconForLanguage returns the flag of the region of the given language tag,
// e.g. "gb" for en-GB. If the tag has no explicit region, the most likely one
// is used ("de" for German). Tags without a determinable region or with a
// region that has no flag (such as es-419) yield an error.
func IconForLanguage(tag language.Tag) (fyne.Resource, error) {
	region, confidence := tag.Region()
	if confidence == language.No {
		return nil, fmt.Errorf("no region for language tag %q: %w", tag, fs.ErrNotExist)
	}
	return Icon(region.String())
}

// MustIconForLanguage is like IconForLanguage but logs the error via
// fyne.LogError and returns theme.ErrorIcon() on failure.
func MustIconForLanguage(tag language.Tag) fyne.Resource {
	icon, err := IconForLanguage(tag)
	if err != nil {
		fyne.LogError("could not load flag icon", err)
		return theme.ErrorIcon()
	}
	return icon
}
