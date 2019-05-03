package canvas

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	pixelDelta = 4 // 4 bytes per pixel for RGBA
)

// Canvas represents a Window with a drawable layer
type Canvas struct {
	Window       *sdl.Window
	Renderer     *sdl.Renderer
	Texture      *sdl.Texture
	Width        int
	Height       int
	CanvasBuffer []byte
}

// Color represents an RGB value for a pixel
type Color struct {
	r, g, b, a byte
}

// Destroy cleans up resources for the Canvas
func (c *Canvas) Destroy() {
	c.Texture.Destroy()
	c.Renderer.Destroy()
	c.Window.Destroy()
	sdl.Quit()
}

// New creates a new Canvas object
func New(title string, width, height int) (*Canvas, error) {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED) // TODO allow toggle to RENDERER_SOFTWARE
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(width), int32(height))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	pixels := make([]byte, width*height*pixelDelta)

	return &Canvas{
		Window:       window,
		Renderer:     renderer,
		Texture:      tex,
		Width:        width,
		Height:       height,
		CanvasBuffer: pixels,
	}, nil
}

// Update repaints the CanvasBuffer
func (c *Canvas) Update() {
	c.Texture.Update(nil, c.CanvasBuffer, c.Width*pixelDelta)
	c.Renderer.Copy(c.Texture, nil, nil)
	c.Renderer.Present()
}

// MakeFullScreen to toggle to a full screen window
func (c *Canvas) MakeFullScreen() {
	c.Window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
}

// SetBackground sets the background color for the Canvas
func (c *Canvas) SetBackground(r, g, b, a byte) {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.setPixel(x, y, Color{r, g, b, a}, c.CanvasBuffer)
		}
	}
}

func (c *Canvas) clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func (c *Canvas) setPixel(x, y int, color Color, pixels []byte) {
	idx := (y*c.Width + x) * pixelDelta

	if idx < len(pixels)-4 && idx >= 0 {
		pixels[idx] = color.r
		pixels[idx+1] = color.g
		pixels[idx+2] = color.b
		pixels[idx+3] = color.a
	}
	// TODO Error case
}
