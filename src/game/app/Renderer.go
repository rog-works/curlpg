package app

import (
	"fmt"
	. "../../graphics"
)

func Renderer(g *Graphics) {
	screen := g.GetScreen()
	lines := ""
	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			lines += string(screen.PixelAt(x, y))
		}
		lines += "\n"
	}
	fmt.Printf(lines)
}