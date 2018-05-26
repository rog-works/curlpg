package scenes

func Factory(key string) Scene {
	if (key == "game") {
		return &GameScene{}
	}
	return nil
}