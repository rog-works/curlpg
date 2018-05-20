package game

import (
	"./app"
	"../graphics"
)

func main() {
	a := app.New()
	a.Run()
	a.Render(graphics.NewGraphics())
}
