package app

import . "../../graphics"

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run() {

}

func (app *Application) Render(g *Graphics) {
	Renderer(g)
}
