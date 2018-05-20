package graphics

type Graphics struct {
	screen *Image
}

func NewGraphics() *Graphics {
	return &Graphics{NewImage(40, 40)}
}

func (g *Graphics) DrawPixel(x, y int, pixel Pixel) {
	g.screen.SetPixel(x, y, pixel)
}

func (g *Graphics) DrawImage(x, y int, image *Image) {
	for offsetX := 0; offsetX < image.Width; offsetX++ {
		for offsetY := 0; offsetY < image.Height; offsetY++ {
			g.DrawPixel(x + offsetX, y + offsetY, image.PixelAt(x, y))
		}
	}
}

func (g *Graphics) DrawChar(x, y int, char rune, ) {
	g.DrawImage(x, y, NewImageFromChar(char))
}

func (g *Graphics) DrawString(x, y int, str string) {
	for _, char := range str {
		g.DrawChar(x, y, char)
	}
}

func (g *Graphics) GetScreen() *Image {
	screen := NewImage(g.screen.Width, g.screen.Height)
	for x := 0; x < g.screen.Width; x++ {
		for y := 0; y < g.screen.Height; y++ {
			screen.SetPixel(x, y, g.screen.PixelAt(x, y))
		}
	}
	return screen;
}
