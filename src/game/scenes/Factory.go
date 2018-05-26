package scenes

func Factory(key string) Scene {
	if (key == "boot") {
		return &BootScene{}
	}
	return nil
}