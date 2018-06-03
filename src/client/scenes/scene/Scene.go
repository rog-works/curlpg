package scene

import "../../graphics"

type Scene interface {
	Run() bool
	Draw(g *graphics.Graphics)
}
