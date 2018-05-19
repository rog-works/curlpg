package main

import "math"

type Screen struct {
	bytes []byte
}

func (s *Screen) Init(width, height int) {
	s.bytes = make(byte, width * height)
}

type Graphics struct {
	screen Screen
}

func (g *Graphics) DrawString(str string) {
	
}

func (g *Graphics) DrawImage(image *Image, offsetX, offsetY int) {
	for var x := 0; x < image.Width; x++ {
		for var y := 0; y < image.Height; y++ {
			if x + offsetX < 0 || y + offsetY < 0 {
				continue
			}
			if image.Width < x + offsetX || image.Height < y + offsetY {
				continue
			}
			g.DrawPixel(image.pixelAt(x, y), x + offsetX, y + offsetY)
		}
	}
}
