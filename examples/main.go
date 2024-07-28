package main

import (
	"fmt"
	_ "image/png"
	"time"

	paintbrush "github.com/jordanella/go-ansi-paintbrush"
)

func main() {
	// Create a new Canvas instance
	canvas := paintbrush.New()

	// Load an image
	err := canvas.LoadImage("examples/norman.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	canvas.Width = 150 // Alternatively, canvas.SetWidth(150)

	// Set threads to 10 (default is 4)
	canvas.Threads = 10 // Alternatively, canvas.SetThreads(10)

	// Add more characters and adjust weights as desired
	var weights = map[rune]float64{
		'': .95,
		'': .95,
		'▁': .9,
		'▂': .9,
		'▃': .9,
		'▄': .9,
		'▅': .9,
		'▆': .85,
		'█': .85,
		'▊': .95,
		'▋': .95,
		'▌': .95,
		'▍': .95,
		'▎': .95,
		'▏': .95,
		'●': .95,
		'◀': .95,
		'▲': .95,
		'▶': .95,
		'▼': .9,
		'○': .8,
		'◉': .95,
		'◧': .9,
		'◨': .9,
		'◩': .9,
		'◪': .9,
	}

	//canvas.AddForbiddenCharacter("M")
	//canvas.AddForbiddenCharacter("@")
	//canvas.AddForbiddenCharacter("#")

	canvas.Weights = weights

	// You can render asynchronously and monitor progress if desired
	canvas.StartPainting()
	for canvas.GetProgress() < 1.0 {
		fmt.Printf("Rendering progress: %.2f%%\r", canvas.GetProgress()*100)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\r%s", canvas.Result)

	// You can also get the C-style string or Bash command if needed:
	//fmt.Println(canvas.GetResultC())
	//fmt.Println(canvas.GetResultBash())
}
