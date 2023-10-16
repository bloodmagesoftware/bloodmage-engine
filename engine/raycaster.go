package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cell_size = 8
)

func RenderViewport() {
	rect := sdl.Rect{
		int32(P.X), int32(P.Y),
		pixel_scale, pixel_scale,
	}
	renderer.SetDrawColor(100, 0, 0, 255)
	renderer.FillRect(&rect)
}
