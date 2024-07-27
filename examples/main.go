package main

import (
	"fmt"
	"time"

	paintbrush "github.com/jordanella/go-ansi-paintbrush"
)

func main() {
	// Create a new AnsiArt instance
	aa := paintbrush.New()

	//aa.AddForbiddenCharacter("M")
	//aa.AddForbiddenCharacter("@")
	//aa.AddForbiddenCharacter("#")

	// Load an image
	err := aa.LoadImage("examples/norman.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	aa.SetWidth(150)

	// Set threads to 10 (default is 4)
	aa.SetThreads(10)

	// You can render asynchronously and monitor progress if desired
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
