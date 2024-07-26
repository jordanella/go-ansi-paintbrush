package paintbrush

type CharacterWeight struct {
	Char   rune
	Weight float64
}

// Add more characters and adjust weights as desired
var CharacterWeights = []CharacterWeight{
	{'', .95},
	{'', .95},
	{'▁', .9},
	{'▂', .9},
	{'▃', .9},
	{'▄', .9},
	{'▅', .9},
	{'▅', .9},
	{'▆', .85},
	{'█', .85},
	{'▊', .95},
	{'▋', .95},
	{'▌', .95},
	{'▍', .95},
	{'▎', .95},
	{'▏', .95},
	{'●', .9},
	{'◀', .95},
	{'▲', .95},
	{'▶', .95},
	{'▼', .9},
	{'○', .8},
	{'◉', .95},
	{'◧', .9},
	{'◨', .9},
	{'◩', .9},
	{'◪', .9},
}

var weightMap map[rune]float64

func init() {
	weightMap = make(map[rune]float64)
	for _, cw := range CharacterWeights {
		weightMap[cw.Char] = cw.Weight
	}
}
