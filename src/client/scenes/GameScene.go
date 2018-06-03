package scenes

import (
	"../graphics"
	"./scene"
	"./game"
)

type states int

const (
	Init states = iota
	Main
	Release
)

type GameScene struct {
	state states
	scenes []scene.Scene
	factory *game.Factory
}

func New() scene.Scene {
	return &GameScene{}
}

func (gs *GameScene) Push(key string) {
	gs.scenes = append(gs.scenes, gs.factory.Create(key))
}

func (gs *GameScene) Run() bool {
	if (gs.state == Init) {
		gs.init()
	} else if (gs.state == Main) {
		gs.main()
	} else if (gs.state == Release) {
		gs.release()
	}
	return true
}

func (gs *GameScene) Draw(g *graphics.Graphics) {
	for _, scene := range gs.scenes {
		scene.Draw(g)
	}
}

func (gs *GameScene) setState(next states) {
	gs.state = next
}

func (gs *GameScene) init() {
	gs.factory = game.NewFactory(gs)
	gs.Push("title")
	gs.setState(Main)
}

func (gs *GameScene) main() {
	for i := len(gs.scenes) - 1; i >= 0; i-- {
		if !gs.scenes[i].Run() {
			gs.scenes = gs.scenes[:i]
		}
	}

	if len(gs.scenes) == 0 {
		gs.setState(Release)
	}
}

func (gs *GameScene) release() {

}
