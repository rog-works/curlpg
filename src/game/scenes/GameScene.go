package scenes

import "../graphics"

type states int

const (
	Init states = iota
	Main
	Release
)

type GameScene struct {
	state states
}

func (s *GameScene) Run() {
	if (s.state == Init) {
		s.init()
	} else if (s.state == Main) {
		s.main()
	} else if (s.state == Release) {
		s.release()
	}
}

func (s *GameScene) Draw(g *graphics.Graphics) {

}
