package paintbrush

import (
	"sort"
	"sync"
)

// StartPainting begins the asynchronous painting process.
func (c *Canvas) StartPainting() {
	go func() {
		c.Paint()
		c.Progress = 1
	}()
}

// GetProgress returns the current progress of the painting process.
func (c *Canvas) GetProgress() float32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Progress
}

func (c *Canvas) renderWorker(wg *sync.WaitGroup, taskChan <-chan Task, resultChan chan<- TaskResult, imgCharWidth, imgCharHeight float64) {
	defer wg.Done()

	for task := range taskChan {
		result := c.processTask(task, imgCharWidth, imgCharHeight)
		resultChan <- result
	}
}

// Paint performs the synchronous painting process.
func (c *Canvas) Paint() {
	if len(c.Font.Glyphs) == 0 {
		fontBytes, err := EmbeddedFonts.ReadFile(FiraMonoRegular)
		if err != nil {
			return
		}
		err = c.SetFont(fontBytes)
		if err != nil {
			return
		}
	}

	c.Result = ""
	c.ResultRGBABytes = nil
	c.ResultC = ""
	c.ResultBash = ""

	fHeight := float64(int(float64(c.Image.Bounds().Dy()) * float64(c.Width) / float64(c.Image.Bounds().Dx()) / c.Font.Aspect))
	imgCharWidth := float64(int((float64(c.Image.Bounds().Dx())/float64(c.Width))*float64(c.GlyphWidth))) / float64(c.GlyphWidth)
	imgCharHeight := float64(c.Image.Bounds().Dy()) / fHeight
	height := int(fHeight)

	c.ResultRGBAWidth = c.Width * c.Font.GlyphWidth
	c.ResultRGBAHeight = height * c.Font.GlyphHeight
	c.ResultRGBABytes = make([]byte, c.ResultRGBAWidth*c.ResultRGBAHeight*4)

	tasks := make([]Task, 0, c.Width*height)
	for charY := 0; charY < height; charY++ {
		for charX := 0; charX < c.Width; charX++ {
			tasks = append(tasks, Task{CharX: charX, CharY: charY})
		}
	}

	sort.Slice(tasks, func(i, j int) bool {
		distFunc := func(t Task) float64 {
			dx := float64(t.CharX) - float64(c.Width)/2
			dy := float64(t.CharY) - fHeight/2
			return dx*dx/c.Font.Aspect + dy*dy*c.Font.Aspect
		}
		di, dj := distFunc(tasks[i]), distFunc(tasks[j])
		if di == dj {
			if tasks[i].CharX == tasks[j].CharX {
				return tasks[i].CharY < tasks[j].CharY
			}
			return tasks[i].CharX < tasks[j].CharX
		}
		return di > dj
	})

	taskResults := make([]TaskResult, len(tasks))
	var wg sync.WaitGroup
	taskChan := make(chan Task, len(tasks))
	resultChan := make(chan TaskResult, len(tasks))

	// Start worker goroutines
	for i := 0; i < c.Threads; i++ {
		wg.Add(1)
		go c.renderWorker(&wg, taskChan, resultChan, imgCharWidth, imgCharHeight)
	}

	// Feed tasks to workers
	go func() {
		for _, task := range tasks {
			taskChan <- task
		}
		close(taskChan)
	}()

	// Collect results
	go func() {
		for i := 0; i < len(tasks); i++ {
			result := <-resultChan
			taskResults[result.CharY*c.Width+result.CharX] = result
			c.Progress = float32(i+1) / float32(len(tasks))
		}
		close(resultChan)
	}()

	wg.Wait()

	// Process results
	c.processResults(taskResults, height)
}
