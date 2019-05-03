package canvas

import (
	"image/png"
	"os"
)

// Image represents a loaded image asset
type Image struct {
	Width  int
	Height int
	Pixels []byte
}

// LoadPNG loads a PNG image from a file path. Can be relative to this file or absolute
func LoadPNG(path string) (*Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	pixels := make([]byte, width*height*4) // for RGBA

	idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[idx] = byte(r / 256)
			idx++
			pixels[idx] = byte(g / 256)
			idx++
			pixels[idx] = byte(b / 256)
			idx++
			pixels[idx] = byte(a / 256)
			idx++
		}
	}

	return &Image{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}, nil
}
