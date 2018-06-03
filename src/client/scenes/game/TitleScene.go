package game

import (
	"../../graphics"
	"../scene"
)

type states int

const (
	Init states = iota
	Main
	Release
)

type TitleScene struct {
	sm scene.Manager
	state states
}

func (s *TitleScene) Run() bool {
	return true
}

func (s *TitleScene) Draw(g *graphics.Graphics) {

}
