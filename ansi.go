package paintbrush

import (
	"image"
	"sync"
)

type AnsiArtInterface interface {
	LoadTTF(data []byte) error
	LoadImage(img image.Image)
	Render()
	StartRender(nThreads int)
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

type AnsiArt struct {
	Font   Font
	Image  image.Image
	width  int
	Height int

	resultRaw        string
	resultRGBABytes  []byte
	resultRGBAWidth  int
	resultRGBAHeight int
	resultC          string
	resultBash       string

	workerCount int
	progress    float32
	mu          sync.Mutex

	forbiddenCharacters map[string]struct{}

	aspectRatio float64
	glyphWidth  int
	glyphHeight int
	runeStart   int
	runeLimit   int
}

func New() AnsiArtInterface {
	return &AnsiArt{
		aspectRatio:         1,
		glyphWidth:          7,
		glyphHeight:         14,
		runeStart:           32,
		runeLimit:           95,
		forbiddenCharacters: make(map[string]struct{}),
	}
}

func (aa *AnsiArt) SetWidth(width int) {
	aa.width = width
}

func (aa *AnsiArt) GetResultBash() string {
	return aa.resultBash
}
func (aa *AnsiArt) GetResultRaw() string {
	return aa.resultRaw
}

func (aa *AnsiArt) GetResultC() string {
	return aa.resultC
}

func (aa *AnsiArt) GetResultRGBABytes() []byte {
	return aa.resultRGBABytes
}

func (aa *AnsiArt) GetResultRGBADimensions() (width, height int) {
	return aa.resultRGBAWidth, aa.resultRGBAHeight
}
func (aa *AnsiArt) AddForbiddenCharacter(char string) {
	aa.forbiddenCharacters[char] = struct{}{}
}

func (aa *AnsiArt) RemoveForbiddenCharacter(char string) {
	delete(aa.forbiddenCharacters, char)
}

func (aa *AnsiArt) ClearForbiddenCharacters() {
	aa.forbiddenCharacters = make(map[string]struct{})
}

func (aa *AnsiArt) GetForbiddenCharacters() []string {
	chars := make([]string, 0, len(aa.forbiddenCharacters))
	for char := range aa.forbiddenCharacters {
		chars = append(chars, char)
	}
	return chars
}

func (aa *AnsiArt) IsForbiddenCharacter(char string) bool {
	_, forbidden := aa.forbiddenCharacters[char]
	return forbidden
}

func (aa *AnsiArt) SetAspectRatio(ratio float64) {
	aa.aspectRatio = ratio
}

func (aa *AnsiArt) GetAspectRatio() float64 {
	return aa.aspectRatio
}

func (aa *AnsiArt) SetGlyphDimensions(width, height int) {
	aa.glyphWidth = width
	aa.glyphHeight = height
}

func (aa *AnsiArt) GetGlyphDimensions() (width, height int) {
	return aa.glyphWidth, aa.glyphHeight
}

func (aa *AnsiArt) SetRuneLimits(start, end int) {
	aa.runeStart = start
	aa.runeLimit = end
}

func (aa *AnsiArt) GetRuneLimits() (int, int) {
	return aa.runeStart, aa.runeLimit
}
