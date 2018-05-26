package scenes

import "../graphics"

type Scene interface {
	Run()
	Draw(g *graphics.Graphics)
}
