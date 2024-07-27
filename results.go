package paintbrush

import (
	"strings"
)

func (aa *AnsiArt) processResults(results []TaskResult, height int) {

	resultIdx := make([][]*TaskResult, height)
	for i := range resultIdx {
		resultIdx[i] = make([]*TaskResult, aa.width)
	}

	for i := range results {
		result := &results[i]
		resultIdx[result.CharY][result.CharX] = result
	}

	var sb strings.Builder
	lastBg := "\033[0m" // Reset background
	lastFg := "\033[0m" // Reset foreground

	for charY := 0; charY < height; charY++ {
		for charX := 0; charX < aa.width; charX++ {
			result := resultIdx[charY][charX]
			if result == nil || result.Glyph == nil {
				sb.WriteString(lastBg + lastFg + " ")
				continue
			}

			var newBg string
			if result.Bg.A < 0.5 {
				newBg = "\033[0m"
			} else {
				newBg = result.Bg.ToPixel().AnsiBg()
			}
			if newBg != lastBg {
				sb.WriteString(newBg)
				lastBg = newBg
			}

			newFg := result.Fg.ToPixel().AnsiFg()
			if newFg != lastFg {
				sb.WriteString(newFg)
				lastFg = newFg
			}
			sb.WriteString(result.Glyph.UTF8)
		}
		// Reset colors at the end of each line
		sb.WriteString("\033[0m\n")
		lastBg = "\033[0m"
		lastFg = "\033[0m"
	}

	// Remove empty newlines at the end
	aa.resultRaw = sb.String()
	for strings.HasSuffix(aa.resultRaw, "\n\n") {
		aa.resultRaw = aa.resultRaw[:len(aa.resultRaw)-1]
	}

	// Generate C string
	aa.resultC = strings.ReplaceAll(aa.resultRaw, "\033", "\\033")
	aa.resultC = strings.ReplaceAll(aa.resultC, "\n", "\\n")
	aa.resultC = strings.ReplaceAll(aa.resultC, "\"", "\\\"")
	aa.resultC = "char kAnsiArt[] = \"" + aa.resultC + "\""

	// Generate Bash string
	aa.resultBash = strings.ReplaceAll(aa.resultRaw, "\\", "\\\\")
	aa.resultBash = strings.ReplaceAll(aa.resultBash, "\033", "\\e")
	aa.resultBash = strings.ReplaceAll(aa.resultBash, "\n", "\\n")
	aa.resultBash = strings.ReplaceAll(aa.resultBash, "'", "\\x27")
	aa.resultBash = "echo -ne '" + aa.resultBash + "'"
}
