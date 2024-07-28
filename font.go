package paintbrush

import (
	"fmt"
	"image"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var FiraMonoRegular = "assets/FiraMono-Regular.ttf"
var FiraMonoBold = "assets/FiraMono-Bold.ttf"

type Font struct {
	GlyphHeight int
	GlyphWidth  int
	Aspect      float64
	Glyphs      []Glyph
}

type Glyph struct {
	Unicode int
	UTF8    string
	Pixels  []uint8
}

func (aa *AnsiArt) SetFont(data []byte) error {
	f, err := truetype.Parse(data)
	if err != nil {
		return err
	}

	// Set fixed glyph dimensions
	aa.Font.GlyphWidth = aa.glyphWidth // You can adjust these values
	aa.Font.GlyphHeight = aa.glyphHeight
	aa.Font.Aspect = (float64(aa.Font.GlyphHeight) / (float64(aa.Font.GlyphWidth))) * aa.aspectRatio

	// Set font size and DPI
	opts := truetype.Options{
		Size:    float64(aa.Font.GlyphHeight), // Use glyph height as font size
		DPI:     72,
		Hinting: font.HintingFull,
	}

	face := truetype.NewFace(f, &opts)

	aa.Font.Glyphs = make([]Glyph, 0)
	for r := rune(aa.runeStart); r < rune(aa.runeLimit); r++ {
		glyph, err := aa.generateGlyph(face, r)
		if err != nil {
			fmt.Printf("Error generating glyph for rune %d: %v\n", r, err)
			continue
		}
		aa.Font.Glyphs = append(aa.Font.Glyphs, glyph)
	}

	for _, cw := range CharacterWeights {
		glyph, err := aa.generateGlyph(face, cw.Char)
		if err != nil {
			fmt.Printf("Error generating glyph for rune %d: %v\n", cw.Char, err)
			continue
		}
		aa.Font.Glyphs = append(aa.Font.Glyphs, glyph)
	}

	return nil
}

func (aa *AnsiArt) generateGlyph(face font.Face, r rune) (Glyph, error) {
	// Create an image to draw the glyph
	img := image.NewGray(image.Rect(0, 0, aa.Font.GlyphWidth, aa.Font.GlyphHeight))
	d := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
	}

	// Get glyph metrics
	advance, ok := face.GlyphAdvance(r)
	if !ok {
		return Glyph{}, fmt.Errorf("glyph not found for rune %v", r)
	}

	// Calculate position to center the glyph
	x := (fixed.I(aa.Font.GlyphWidth) - advance) / 2
	y := fixed.I(aa.Font.GlyphHeight * 4 / 5) // Adjust this value to vertically center the glyph

	// Draw the glyph
	d.Dot = fixed.Point26_6{X: x, Y: y}
	d.DrawString(string(r))

	// Convert image to pixel array
	pixels := make([]uint8, aa.Font.GlyphWidth*aa.Font.GlyphHeight)
	for y := 0; y < aa.Font.GlyphHeight; y++ {
		for x := 0; x < aa.Font.GlyphWidth; x++ {
			pixels[y*aa.Font.GlyphWidth+x] = img.GrayAt(x, y).Y
		}
	}

	return Glyph{
		Unicode: int(r),
		UTF8:    string(r),
		Pixels:  pixels,
	}, nil
}
