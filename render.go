package paintbrush

import (
	"sort"
	"sync"
)

func (aa *AnsiArt) StartRender(nThreads int) {
	aa.workerCount = nThreads
	go func() {
		aa.Render()
		aa.progress = 1
	}()
}

func (aa *AnsiArt) GetRenderProgress() float32 {
	aa.mu.Lock()
	defer aa.mu.Unlock()
	return aa.progress
}

func (aa *AnsiArt) renderWorker(wg *sync.WaitGroup, taskChan <-chan Task, resultChan chan<- TaskResult, imgCharWidth, imgCharHeight float64) {
	defer wg.Done()

	for task := range taskChan {
		result := aa.processTask(task, imgCharWidth, imgCharHeight)
		resultChan <- result
	}
}

func (aa *AnsiArt) Render() {
	aa.resultRaw = ""
	aa.resultRGBABytes = nil
	aa.resultC = ""
	aa.resultBash = ""

	fheight := float64(int(float64(aa.Image.Bounds().Dy()) * float64(aa.width) / float64(aa.Image.Bounds().Dx()) / aa.Font.Aspect))
	imgCharWidth := float64(int((float64(aa.Image.Bounds().Dx())/float64(aa.width))*float64(aa.glyphWidth))) / float64(aa.glyphWidth)
	imgCharHeight := float64(aa.Image.Bounds().Dy()) / fheight
	height := int(fheight)

	aa.resultRGBAWidth = aa.width * aa.Font.GlyphWidth
	aa.resultRGBAHeight = height * aa.Font.GlyphHeight
	aa.resultRGBABytes = make([]byte, aa.resultRGBAWidth*aa.resultRGBAHeight*4)

	tasks := make([]Task, 0, aa.width*height)
	for charY := 0; charY < height; charY++ {
		for charX := 0; charX < aa.width; charX++ {
			tasks = append(tasks, Task{CharX: charX, CharY: charY})
		}
	}

	sort.Slice(tasks, func(i, j int) bool {
		distFunc := func(t Task) float64 {
			dx := float64(t.CharX) - float64(aa.width)/2
			dy := float64(t.CharY) - fheight/2
			return dx*dx/aa.Font.Aspect + dy*dy*aa.Font.Aspect
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
	for i := 0; i < aa.workerCount; i++ {
		wg.Add(1)
		go aa.renderWorker(&wg, taskChan, resultChan, imgCharWidth, imgCharHeight)
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
			taskResults[result.CharY*aa.width+result.CharX] = result
			aa.progress = float32(i+1) / float32(len(tasks))
		}
		close(resultChan)
	}()

	wg.Wait()

	// Process results
	aa.processResults(taskResults, height)
}
