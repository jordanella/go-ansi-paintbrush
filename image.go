package paintbrush

import (
	"fmt"
	"image"
	_ "image/png" // PNG Support
	"math"
	"os"
)

func (aa *AnsiArt) LoadImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return err
	}

	aa.SetImage(img)
	return nil
}

func (aa *AnsiArt) SetImage(img image.Image) {
	aa.Image = img
}

func (aa *AnsiArt) readImageColor(x, y float64) Vec4 {
	if x >= float64(aa.Image.Bounds().Dx()) || x < 0 || y >= float64(aa.Image.Bounds().Dy()) || y < 0 {
		return Vec4{}
	}
	r, g, b, a := aa.Image.At(int(math.Round(x)), int(math.Round(y))).RGBA()
	return Vec4{
		R: float64(r) / 65535.0,
		G: float64(g) / 65535.0,
		B: float64(b) / 65535.0,
		A: float64(a) / 65535.0,
	}
}
