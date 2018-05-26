package app

import (
	"../graphics"
	"../scenes"
)

type Application struct {
	screen *graphics.Image
	g *graphics.Graphics
	scene scenes.Scene
}

func New() *Application {
	screen := graphics.NewImage(20, 5)
	g := graphics.NewGraphics(screen)
	scene := scenes.Factory("boot")
	return &Application{screen, g, scene}
}

func (app *Application) Run() {
	app.main()
	app.render()
	app.flip()
}

func (app *Application) main() {
	app.scene.Run()
}

func (app *Application) render() {
	app.scene.Draw(app.g)
}

func (app *Application) flip() {
	Renderer(app.screen)
}

