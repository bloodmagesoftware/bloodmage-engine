// Bloodmage Engine
// Copyright (C) 2024 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

package topdown

import (
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
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
