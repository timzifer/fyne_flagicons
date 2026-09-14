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

// by ISO 3166-1 alpha-2 code (case-insensitive) or generated constant
de := widget.NewIcon(fyne_flagicons.MustIcon(fyne_flagicons.FlagDE))
eng := widget.NewIcon(fyne_flagicons.MustIcon("gb-eng"))

// by language tag: uses the tag's region, e.g. en-GB -> gb, de -> de
gb := widget.NewIcon(fyne_flagicons.MustIconForLanguage(language.BritishEnglish))

// raw SVG or a rasterized PNG (longer edge in pixels)
svg, _ := fyne_flagicons.Source("fr")
png, _ := fyne_flagicons.PNG("fr", 64)
```

`Icon` / `IconForLanguage` return an error wrapping `fs.ErrNotExist` for unknown
codes; the `Must*` variants log via `fyne.LogError` and return
`theme.ErrorIcon()` instead. `Names()` lists all codes.

Built on [fyne_iconkit](https://github.com/timzifer/fyne_iconkit). After
updating the SVGs, run `go generate ./...` to refresh the `Flag*` constants.

## License

Code: [MIT](LICENSE). Flag SVGs: MIT, © Panayiotis Lipiridis – see
[LICENSE-flag-icons](LICENSE-flag-icons).
