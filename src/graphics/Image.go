package graphics

type Image struct {
	pixels []Pixel
	Width  int
	Height int
}

func NewImage(width, height int) *Image {
	image := &Image{make([]Pixel, width * height), width, height}
	image.Clear()
	return image
}

func NewImageFromChar(char rune) *Image {
	image := &Image{make([]Pixel, 1), 1, 1}
	image.SetPixel(0, 0, Pixel(char))
	return image
}

func (image *Image) Clear() {
	for i := 0; i < image.size(); i++ {
		image.pixels[i] = Pixel([]rune(" ")[0])
	}
}

func (image *Image) SetPixel(x, y int, pixel Pixel) {
	if (image.isInside(x, y)) {
		image.pixels[image.toIndex(x, y)] = pixel
	}
}

func (image *Image) PixelAt(x, y int) Pixel {
	if (image.isInside(x, y)) {
		return image.pixels[image.toIndex(x, y)]
	} else {
		return 0
	}
}

func (image *Image) toIndex(x, y int) int {
	return y * image.Width + x;
}

func (image *Image) size() int {
	return image.toIndex(image.Width, image.Height)
}

func (image *Image) isInside(x, y int) bool {
	return 0 <= x && x < image.Width &&
		   0 <= y && y < image.Height
}
