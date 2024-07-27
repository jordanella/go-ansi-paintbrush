# ANSI Paintbrush

A Go fork of the [C++ ANSI Art rendering library](https://github.com/mafik/ansi-art) originally created by [Marek Rogalski](https://github.com/mafik).

ANSI Paintbrush allows you to convert images into colorful ASCII art using ANSI escape codes. It provides a simple interface for loading images, rendering them as ASCII art, and outputting the result in various formats.

## Example

![Example Output](docs/norman.png)

## Features

- Convert images to ASCII art with ANSI color codes
- Load custom TTF fonts for character selection
- Adjustable output width
- Multi-threaded rendering for improved performance
- Multiple output formats (raw string, C-style string, Bash command)
- Weighting and adding specific characters
- Ability to exclude specific characters entirely

## Installation

```
go get github.com/jordanella/go-ansi-paintbrush
```

## Basic Usage

```
package main

import (
    "fmt"

    "github.com/jordanella/go-ansi-paintbrush"
)

func main() {
    // Create a new AnsiArt instance
    aa := paintbrush.New()

    // Load an image
    err := aa.LoadImage("examples/norman.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start the rendering process
	aa.Render()

    // Print the result
    fmt.Printf("\r%s", aa.GetResultRaw())
}
```

## Planned Features

- Configurable height constraint
- Advanced sizing options (stretch, bottleneck, crop, etc.)
- Command-line argument handling
- Render image to file

## Interface Reference

```
type AnsiArtInterface interface {
    LoadFont(path string) error
    SetFont(data []byte) error
    LoadImage(path string) error
    SetImage(img image.Image)
    Render()
    SetThreads(int)
    StartRender()
    GetRenderProgress() float32
    GetResultRaw() string
    GetResultC() string
    GetResultBash() string
    GetResultRGBABytes() []byte
    GetResultRGBADimensions() (width, height int)
    SetWidth(int)
    AddForbiddenCharacter(string)
    RemoveForbiddenCharacter(string)
    ClearForbiddenCharacters()
    GetForbiddenCharacters() []string
    IsForbiddenCharacter(string) bool
    SetAspectRatio(float64)
    GetAspectRatio() float64
    SetGlyphDimensions(width, height int)
    GetGlyphDimensions() (width, height int)
    SetRuneLimits(start, end int)
    GetRuneLimits() (start, end int)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

This project is a Go fork of the original [C++ ANSI Art library](https://github.com/mafik/ansi-art) created by Marek Rogalski. I am very grateful for their work, which served as the foundation for this Go implementation.
