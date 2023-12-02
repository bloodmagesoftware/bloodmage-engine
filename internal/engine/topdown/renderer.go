package topdown

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/level"
	"github.com/veandco/go-sdl2/sdl"
)

func minInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// RenderViewport renders a tow down view of the current level with the player in the center.
func RenderViewport() {
	tileSize := minInt32(core.Width(), core.Height()) / 10

	// draw floor
	rect := sdl.Rect{X: 0, Y: 0, W: tileSize, H: tileSize}
	for y := 0; y != level.Height(); y++ {
		rect.Y = int32(y)*tileSize - int32(core.P.Y*float32(tileSize)) + int32(core.Height())/2

		if rect.Y > int32(core.Height()) {
			continue
		}

		if rect.Y+rect.H < 0 {
			continue
		}

		for x := 0; x != level.Width(); x++ {
			rect.X = int32(x)*tileSize - int32(core.P.X*float32(tileSize)) + int32(core.Width())/2

			if rect.X > int32(core.Width()) {
				continue
			}

			if rect.X+rect.W < 0 {
				continue
			}

			if level.Collision(x, y) {
				_ = core.Renderer().SetDrawColor(128, 128, 128, 255)
			} else {
				_ = core.Renderer().SetDrawColor(8, 8, 8, 255)
			}

			_ = core.Renderer().FillRect(&rect)
		}
	}

	// draw player
	rect.X = int32(core.Width())/2 - tileSize/2
	rect.Y = int32(core.Height())/2 - tileSize/2
	_ = core.Renderer().SetDrawColor(255, 255, 255, 255)
	_ = core.Renderer().FillRect(&rect)
}
