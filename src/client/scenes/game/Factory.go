package game

import "../scene"

type Factory struct {
	sm scene.Manager
}

func NewFactory(sm scene.Manager) *Factory {
	return &Factory{sm}
}

func (f *Factory) Create(key string) scene.Scene {
	if (key == "title") {
		return &TitleScene{f.sm, Init}
	}
	return nil
}