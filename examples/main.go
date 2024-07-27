package main

import (
	"fmt"
	"image"
	_ "image/png" // Import this to support PNG images
	"os"
	"time"

	paintbrush "github.com/jordanella/go-ansi-paintbrush"
)

func main() {
	// Create a new AnsiArt instance
	aa := paintbrush.New()

	//aa.AddForbiddenCharacter("M")
	//aa.AddForbiddenCharacter("@")
	//aa.AddForbiddenCharacter("#")

	// Load a ttf font
	fontPath := ""
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		fmt.Printf("Error reading font file: %v\n", err)
		return
	}
	err = aa.LoadTTF(fontData)
	if err != nil {
		fmt.Printf("Error loading TTF: %v\n", err)
		return
	}

	// Load an image
	imagePath := ""
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	aa.LoadImage(img)

	// Set the desired width of the output
	aa.SetWidth(150)

	// Set threads to 10 (default is 4)
	aa.SetThreads(10)

	// Start the rendering process
	aa.StartRender()

	for aa.GetRenderProgress() < 1.0 {
		fmt.Printf("Rendering progress: %.2f%%\r", aa.GetRenderProgress()*100)
		time.Sleep(100 * time.Millisecond)
	}

	result := aa.GetResultRaw()
	fmt.Printf("\r%s", result)

	// You can also get the C-style string or Bash command if needed:
	//fmt.Println(aa.GetResultC())
	//fmt.Println(aa.GetResultBash())
}
