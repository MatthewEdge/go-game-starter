package canvas

// Texture represents a texture saved on the Canvas
type Texture struct {
	Pixels               []byte
	Width, Height, Pitch int
}

// DrawTexture adds a new Texture to the Canvas
func (c *Canvas) DrawTexture(name string, px, py int, tex *Texture) {
	for y := 0; y < tex.Height; y++ {
		for x := 0; x < tex.Width; x++ {
			screenY := y + py
			screenX := x + px

			if screenX >= 0 && screenX < c.Width && screenY >= 0 && screenY < c.Height {
				texIndex := y*tex.Pitch + x*4
				screenIndex := screenY*c.Width*4 + screenX*4

				c.CanvasBuffer[screenIndex] = tex.Pixels[texIndex]     // R
				c.CanvasBuffer[screenIndex+1] = tex.Pixels[texIndex+1] // G
				c.CanvasBuffer[screenIndex+2] = tex.Pixels[texIndex+2] // B
				c.CanvasBuffer[screenIndex+3] = tex.Pixels[texIndex+3] // A
			}
		}
	}
}

// DrawAlphaTexture adds a new Texture to the Canvas using alpha blending (more expensive)
func (c *Canvas) DrawAlphaTexture(name string, px, py int, tex *Texture) {
	for y := 0; y < tex.Height; y++ {
		for x := 0; x < tex.Width; x++ {
			screenY := y + py
			screenX := x + px

			if screenX >= 0 && screenX < c.Width && screenY >= 0 && screenY < c.Height {
				texIndex := y*tex.Pitch + x*4
				screenIndex := screenY*c.Width*4 + screenX*4

				srcR := int(tex.Pixels[texIndex])
				srcG := int(tex.Pixels[texIndex+1])
				srcB := int(tex.Pixels[texIndex+2])
				srcA := int(tex.Pixels[texIndex+3])

				destR := int(c.CanvasBuffer[screenIndex])
				destG := int(c.CanvasBuffer[screenIndex+1])
				destB := int(c.CanvasBuffer[screenIndex+2])

				// Scale by 255 since we're doing RGBA values of 0-255 vs. the typical alpha blending values of 0 - 1
				rstR := (srcR*255 + destR*(255-srcA)) / 255
				rstG := (srcG*255 + destG*(255-srcA)) / 255
				rstB := (srcB*255 + destB*(255-srcA)) / 255

				c.CanvasBuffer[screenIndex] = byte(rstR)   // R
				c.CanvasBuffer[screenIndex+1] = byte(rstG) // G
				c.CanvasBuffer[screenIndex+2] = byte(rstB) // B
				// c.CanvasBuffer[screenIndex+3] = tex.Pixels[texIndex+3] // A
			}
		}
	}
}
