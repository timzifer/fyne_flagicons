// Package fyne_flagicons provides country flag icons for Fyne apps.
//
// Flags are addressed by the generated Flag* variables (e.g. FlagDE,
// FlagGB_ENG), by Lookup for ISO 3166-1 alpha-2 codes from configuration, or
// by ForLanguage for language tags. The SVGs are taken from
// https://github.com/lipis/flag-icons.
package fyne_flagicons

import (
	"embed"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/timzifer/fyne_iconkit"
	"golang.org/x/text/language"
)

//go:generate go run github.com/timzifer/fyne_iconkit/cmd/iconconst -dir flags -type flag -prefix Flag -upper -out flags_gen.go

// flag identifies one embedded flag. It is unexported on purpose: values only
// come from the generated Flag* variables, Lookup or ForLanguage, so every
// flag exists.
type flag struct{ code string }

// String returns the lowercase flag code, e.g. "de" or "gb-eng".
func (f flag) String() string { return f.code }

var (
	//go:embed flags
	flags embed.FS

	set = fyne_iconkit.NewSet(flags, "flags")
)

// Lookup returns the flag for a country code such as "de", "DE" or "gb-eng".
func Lookup(code string) (flag, bool) {
	code = strings.ToLower(code)
	if !set.Has(code) {
		return flag{}, false
	}
	return flag{code}, true
}

// ForLanguage returns the flag of the region of the given language tag, e.g.
// FlagGB for en-GB. Without an explicit region the most likely one is used
// (FlagDE for German). It reports false if no region can be determined or the
// region has no flag (such as es-419).
func ForLanguage(tag language.Tag) (flag, bool) {
	region, confidence := tag.Region()
	if confidence == language.No {
		return flag{}, false
	}
	return Lookup(region.String())
}

// All returns every flag of this package, sorted by code.
func All() []flag {
	codes := set.Names()
	all := make([]flag, len(codes))
	for n, code := range codes {
		all[n] = flag{code}
	}
	return all
}

// Icon returns f as a cached fyne.Resource.
func Icon(f flag) fyne.Resource {
	return set.MustIcon(f.code)
}

// IconForLanguage returns the flag icon for the given language tag, or
// theme.ErrorIcon() if ForLanguage finds no flag.
func IconForLanguage(tag language.Tag) fyne.Resource {
	f, _ := ForLanguage(tag)
	return Icon(f)
}

// Source returns the SVG content of f.
func Source(f flag) []byte {
	src, _ := set.Source(f.code)
	return src
}

// PNG rasterizes f. size is the length of the longer edge in pixels.
func PNG(f flag, size int) ([]byte, error) {
	return set.PNG(f.code, size)
}
