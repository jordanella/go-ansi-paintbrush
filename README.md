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

### Future Plans

I'm always looking to improve ANSI Paintbrush. Some features being considering for future releases include:

- Configurable height constraint
- Advanced sizing options (stretch, bottleneck, crop, etc.)
- Command-line argument handling
- Rendering the result to file

## Installation

```bash
go get github.com/jordanella/go-ansi-paintbrush
```

## Quickstart

```go
package main

import (
    "fmt"
	_ "image/png"

    "github.com/jordanella/go-ansi-paintbrush"
)

func main() {
    // Create a new AnsiArt instance
    canvas := paintbrush.New()

    // Load an image
    err := canvas.LoadImage("examples/norman.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start the rendering process
	canvas.Paint()

    // Print the result
    fmt.Printf("\r%s", canvas.Result)
}
```

Note that it is important to include the appropriate file type support necessary for your project.

```go
import (
    _ "image/png" // PNG support example
)

```

## Canvas Reference

### Type
The Canvas struct is the core of the ANSI Paintbrush library. It contains all the necessary fields for image processing, rendering, and output generation.
```go
type Canvas struct {
	// Input and Rendering Configuration
	Font                Font              // Font used for rendering
	Image               image.Image       // Input image to be processed
	Width               int               // Output width in characters
	Height              int               // Output height in characters
	AspectRatio         float64           // Aspect ratio for output
	GlyphWidth          int               // Width of each glyph
	GlyphHeight         int               // Height of each glyph
	RuneStart           int               // Starting Unicode code point for character selection
	RuneLimit           int               // Ending Unicode code point for character selection
	Threads             int               // Number of threads for parallel processing
	ForbiddenCharacters map[rune]struct{} // Characters to exclude from rendering
	Weights             map[rune]float64  // Custom weights for character selection

    // Output Results
    Result       string         // Raw output string
    ResultC      string         // C-style string output
    ResultBash   string         // Bash command string output
    ResultRGBABytes []byte      // RGBA byte slice of the rendered image
    ResultRGBAWidth int         // Width of the RGBA output
    ResultRGBAHeight int        // Height of the RGBA output

    // Internal State
    Progress     float32        // Current progress of rendering (0.0 to 1.0)
    mu           sync.Mutex     // Mutex for thread-safe operations
}
```

### Methods

The ```Canvas``` struct provides the following methods. For detailed documentation on each method, please refer to the inline comments in the source code or visit the GoDoc documentation.

#### Initialization
```go
New() *Canvas
```

#### Input and Rendering Configuration
```go
LoadFont(path string) error
SetFont(data []byte) error
LoadImage(path string) error
SetImage(img image.Image)
GetImage() image.Image
SetWidth(int)
SetHeight(int)
SetAspectRatio(float64)
GetAspectRatio() float64
SetGlyphDimensions(width, height int)
GetGlyphDimensions() (width, height int)
SetRuneLimits(start, end int)
GetRuneLimits() (start, end int)
SetThreads(int)
AddForbiddenCharacter(rune)
RemoveForbiddenCharacter(rune)
ClearForbiddenCharacters()
GetForbiddenCharacters() []rune
IsForbiddenCharacter(rune) bool
SetWeights(map[rune]float64)
AddWeights(map[rune]float64)
```

#### Rendering Process
```go
Paint()
StartPainting()
GetProgress() float32
```

#### Output Retrieval
```go
GetResult() string
GetResultC() string
GetResultBash() string
GetResultRGBABytes() []byte
GetResultRGBADimensions() (width, height int)
```

## Character Weighting and Custom Characters

The ANSI Paintbrush library allows you to customize the character selection process through a weighting system. Weightings can be leveraged to emphasize certain characters over others or to add entirely new characters to the rendering process. This flexibility allows you to fine-tune the output to achieve the desired aesthetic for your images.

Characters with higher weights (closer to 1.0) are more likely to be chosen during the rendering process. The default weight for all characters is 1.0. Any specific characters will also be added to the pool of available characters for rendering, with the corresponding weights.

Note: Characters with weights set to 0 or negative values will be excluded from the rendering process entirely.

### Setting Weights

You can set weights for characters using the `SetWeights` method, which replaces the entire existing weight map:
```go
weights := map[rune]float64{
    '█': 0.95,
    '▓': 0.90,
    '▒': 0.85,
    '░': 0.80,
    '●': 0.75,
}
canvas.SetWeights(weights)
```

### Adding or Updating Weights

To add new weights or update existing ones without affecting other characters, use the AddWeights method:
```
newWeights := map[rune]float64{
    '♥': 0.9,
    '♦': 0.9,
    '▓': 0.95, // This will update the existing weight for '▓'
}
canvas.AddWeights(newWeights)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

This project is a Go fork of the original [C++ ANSI Art library](https://github.com/mafik/ansi-art) created by Marek Rogalski. I am very grateful for their work, which served as the foundation for this Go implementation.
