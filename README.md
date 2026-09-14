# fyne_flagicons

[![CI](https://github.com/timzifer/fyne_flagicons/actions/workflows/ci.yml/badge.svg)](https://github.com/timzifer/fyne_flagicons/actions/workflows/ci.yml)

Country flag icons for [Fyne](https://fyne.io) apps, embedded into your binary.
The SVGs come from [flag-icons](https://github.com/lipis/flag-icons).
Requires Go 1.21+ and Fyne 2.3+.

## Install

```sh
go get github.com/timzifer/fyne_flagicons
```

## Usage

```go
import (
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/language"

	"github.com/timzifer/fyne_flagicons"
)

// generated variables
de := widget.NewIcon(fyne_flagicons.Icon(fyne_flagicons.FlagDE))
eng := widget.NewIcon(fyne_flagicons.Icon(fyne_flagicons.FlagGB_ENG))

// by language tag: uses the tag's region, e.g. en-GB -> FlagGB, de -> FlagDE
gb := widget.NewIcon(fyne_flagicons.IconForLanguage(language.BritishEnglish))

// codes from configuration or user input (case-insensitive)
if f, ok := fyne_flagicons.Lookup("fr"); ok {
	svg := fyne_flagicons.Source(f)
	png, _ := fyne_flagicons.PNG(f, 64) // longer edge in pixels
}
```

Flags are values of an unexported type that only the generated `Flag*`
variables, `Lookup` and `ForLanguage` produce – arbitrary strings do not
compile, so `Icon` cannot fail. `IconForLanguage` falls back to
`theme.ErrorIcon()` if the tag has no flag (e.g. `es-419`); use `ForLanguage`
to check first. Since the type is unexported, store `fyne.Resource`s or codes
(`f.String()`) rather than flag values in your own structs.

Built on [fyne_iconkit](https://github.com/timzifer/fyne_iconkit). After
updating the SVGs, run `go generate ./...` to refresh the `Flag*` variables.

## License

Code: [MIT](LICENSE). Flag SVGs: MIT, © Panayiotis Lipiridis – see
[LICENSE-flag-icons](LICENSE-flag-icons).
