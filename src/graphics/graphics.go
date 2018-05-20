package graphics

import "math"

type Pixel struct {
	buf byte
}

type Screen struct {
	pixels []Pixel
	Width int
	Height int
}

type Image Screen

func New(width, height int) (*Screen) {
	return &Screen{
		make(Pixel, width * height),
		width,
		height
	}
}

type Graphics struct {
	screen Screen
}

func (g *Graphics) DrawPixel(pixel Pixel, offsetX, offsetY int) {
	if offsetX < 0 || offsetY < 0 {
		continue
	}
	if g.scree.Width < offsetX || g.scree.Height < offsetY {
		continue
	}
	var index := offsetY * g.screen.Width + offsetX
	g.screen[index] = pixel
}

func (g *Graphics) DrawImage(image *Image, offsetX, offsetY int) {
	for var x := 0; x < image.Width; x++ {
		for var y := 0; y < image.Height; y++ {
			g.DrawPixel(image.PixelAt(x, y), x + offsetX, y + offsetY)
		}
	}
}

func (g *Graphics) DrawChar(char string, offsetX, offsetY int) {
	g.DrawPixel(new Pixel(char), offsetX, offsetY)
}

func (g *Graphics) DrawString(str string, offsetX, offsetY) {
	for var i := 0; i < str.Length; i++ {
		g.DrawChar(str[i], offsetX, offsetY)
	}
}

