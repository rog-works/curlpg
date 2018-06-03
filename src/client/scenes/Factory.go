package scenes

import "./scene"

func Factory(key string) scene.Scene {
	if (key == "game") {
		return &GameScene{}
	}
	return nil
}