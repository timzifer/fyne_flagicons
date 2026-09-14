package fyne_flagicons

import (
	"bytes"
	"errors"
	"image/png"
	"io/fs"
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"golang.org/x/text/language"
)

func TestAllFlagsRender(t *testing.T) {
	names := Names()
	if len(names) < 250 {
		t.Fatalf("only %d flags embedded", len(names))
	}
	for _, name := range names {
		data, err := PNG(name, 32)
		if err != nil {
			t.Errorf("PNG(%q): %v", name, err)
			continue
		}
		if _, err := png.Decode(bytes.NewReader(data)); err != nil {
			t.Errorf("PNG(%q) is not decodable: %v", name, err)
		}
	}
}

func TestIcon(t *testing.T) {
	for _, code := range []string{FlagDE, "DE", "de.svg", FlagGB_ENG} {
		res, err := Icon(code)
		if err != nil {
			t.Errorf("Icon(%q): %v", code, err)
			continue
		}
		if !bytes.HasPrefix(res.Content(), []byte("<svg")) {
			t.Errorf("Icon(%q) content is not an SVG", code)
		}
	}
	if _, err := Icon("nope"); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("Icon(nope) error = %v, want fs.ErrNotExist", err)
	}
}

func TestIconForLanguage(t *testing.T) {
	tests := []struct {
		tag  language.Tag
		want string
	}{
		{language.MustParse("de-AT"), FlagAT},
		{language.BritishEnglish, FlagGB},
		{language.German, FlagDE},
		{language.MustParse("pt-BR"), FlagBR},
	}
	for _, tt := range tests {
		res, err := IconForLanguage(tt.tag)
		if err != nil {
			t.Errorf("IconForLanguage(%s): %v", tt.tag, err)
			continue
		}
		if res.Name() != tt.want+".svg" {
			t.Errorf("IconForLanguage(%s) = %s, want %s.svg", tt.tag, res.Name(), tt.want)
		}
	}

	for _, tag := range []language.Tag{language.MustParse("es-419")} {
		if _, err := IconForLanguage(tag); !errors.Is(err, fs.ErrNotExist) {
			t.Errorf("IconForLanguage(%s) error = %v, want fs.ErrNotExist", tag, err)
		}
	}
}

func TestMustIconForLanguageFallback(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	if got := MustIconForLanguage(language.MustParse("es-419")); got != theme.ErrorIcon() {
		t.Errorf("MustIconForLanguage(es-419) = %v, want theme.ErrorIcon()", got)
	}
	if got := MustIcon("nope"); got != theme.ErrorIcon() {
		t.Errorf("MustIcon(nope) = %v, want theme.ErrorIcon()", got)
	}
}
