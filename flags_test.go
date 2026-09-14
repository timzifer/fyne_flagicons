package fyne_flagicons

import (
	"bytes"
	"image/png"
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"golang.org/x/text/language"
)

func TestAllFlagsRender(t *testing.T) {
	all := All()
	if len(all) < 250 {
		t.Fatalf("only %d flags embedded", len(all))
	}
	for _, f := range all {
		data, err := PNG(f, 32)
		if err != nil {
			t.Errorf("PNG(%s): %v", f, err)
			continue
		}
		if _, err := png.Decode(bytes.NewReader(data)); err != nil {
			t.Errorf("PNG(%s) is not decodable: %v", f, err)
		}
	}
}

func TestLookup(t *testing.T) {
	tests := map[string]flag{"de": FlagDE, "DE": FlagDE, "gb-eng": FlagGB_ENG, "GB-ENG": FlagGB_ENG}
	for code, want := range tests {
		if got, ok := Lookup(code); !ok || got != want {
			t.Errorf("Lookup(%q) = %v, %v; want %v", code, got, ok, want)
		}
	}
	for _, code := range []string{"", "nope", "de.svg", "../flags/de"} {
		if _, ok := Lookup(code); ok {
			t.Errorf("Lookup(%q) succeeded", code)
		}
	}
}

func TestForLanguage(t *testing.T) {
	tests := []struct {
		tag  language.Tag
		want flag
	}{
		{language.MustParse("de-AT"), FlagAT},
		{language.BritishEnglish, FlagGB},
		{language.German, FlagDE},
		{language.MustParse("pt-BR"), FlagBR},
	}
	for _, tt := range tests {
		if got, ok := ForLanguage(tt.tag); !ok || got != tt.want {
			t.Errorf("ForLanguage(%s) = %v, %v; want %v", tt.tag, got, ok, tt.want)
		}
	}
	if got, ok := ForLanguage(language.MustParse("es-419")); ok {
		t.Errorf("ForLanguage(es-419) = %v, want no flag", got)
	}
}

func TestIcon(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	res := Icon(FlagDE)
	if res.Name() != "de.svg" || !bytes.HasPrefix(res.Content(), []byte("<svg")) {
		t.Errorf("Icon(FlagDE) = %s %.20s", res.Name(), res.Content())
	}
	if !bytes.Equal(Source(FlagDE), res.Content()) {
		t.Error("Source(FlagDE) differs from icon content")
	}
	if got := IconForLanguage(language.German); got != res {
		t.Errorf("IconForLanguage(de) = %v, want %v", got, res)
	}
	if got := IconForLanguage(language.MustParse("es-419")); got != theme.ErrorIcon() {
		t.Errorf("IconForLanguage(es-419) = %v, want theme.ErrorIcon()", got)
	}

	var zero flag
	if got := Icon(zero); got != theme.ErrorIcon() {
		t.Errorf("Icon(zero value) = %v, want theme.ErrorIcon()", got)
	}
}
