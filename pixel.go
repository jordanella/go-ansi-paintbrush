package paintbrush

import "fmt"

type Pixel struct {
	R, G, B, A uint8
}

func (p Pixel) AnsiColor() string {
	return fmt.Sprintf("2;%d;%d;%d", p.R, p.G, p.B)
}

func (p Pixel) AnsiBg() string {
	return "\033[48;" + p.AnsiColor() + "m"
}

func (p Pixel) AnsiFg() string {
	return "\033[38;" + p.AnsiColor() + "m"
}
