package app

import (
	. "../../graphics"
	. "../scenes"
)

type Application struct {
	screen *Image
	g *Graphics
	scene *Scene
}

func New() *Application {
	screen := NewImage(20, 5)
	g := NewGraphics(screen)
	scene := Scene.Factory('game')
	return &Application{screen, g}
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

