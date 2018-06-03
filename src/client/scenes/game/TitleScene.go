package game

import (
	"../../graphics"
	"../../input"
	"../scene"
)

type states int

const (
	Init states = iota
	Main
	ToGame
	Release
)

type TitleScene struct {
	sm scene.Manager
	state states
}

func (ts *TitleScene) Run() bool {
	if (ts.state == Init) {
		return ts.init()
	} else if (ts.state == Main) {
		return ts.main()
	} else if (ts.state == Release) {
		return ts.release()
	}
	return false
}

func (s *TitleScene) Draw(g *graphics.Graphics) {

}

func (ts *TitleScene) setState(next states) {
	ts.state = next
}

func (ts *TitleScene) init() bool {
	ts.setState(Main)
	return true
}

func (ts *TitleScene) main() bool {
	if input.Pushed(input.Accept) {
		ts.setState(ToGame)
	} else if input.Pushed(input.Cancel) {
		ts.setState(Release)
	}
	return true
}

func (ts *TitleScene) toGame() bool {
	ts.sm.Push("game")
	ts.setState(Release)
	return true
}

func (ts *TitleScene) release() bool {
	return false
}
