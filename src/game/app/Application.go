package app

import . "../../graphics"

type Application struct {
	screen *Image
	g *Graphics
}

func New() *Application {
	screen := NewImage(20, 5)
	g := NewGraphics(screen)
	return &Application{screen, g}
}

func (app *Application) Run() {
	app.main()
	app.render()
	app.flip()
}

func (app *Application) main() {
}

func (app *Application) render() {
	app.g.DrawString(5, 0, "hoge")
	app.g.DrawString(0, 1, "fuga")
}

func (app *Application) flip() {
	Renderer(app.screen)
}

