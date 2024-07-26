package paintbrush

import (
	"image"
	"math"
)

func (aa *AnsiArt) LoadImage(img image.Image) {
	aa.Image = img
	aa.width = img.Bounds().Dx()
	aa.Height = img.Bounds().Dy()
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
