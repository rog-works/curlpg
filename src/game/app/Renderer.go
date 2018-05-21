package app

import (
	"fmt"
	. "../../graphics"
)

func Renderer(screen *Image) {
	lines := ""
	for y := 0; y < screen.Height; y++ {
		for x := 0; x < screen.Width; x++ {
			lines += string(screen.PixelAt(x, y))
		}
		lines += "\n"
	}
	fmt.Printf(lines)
}
