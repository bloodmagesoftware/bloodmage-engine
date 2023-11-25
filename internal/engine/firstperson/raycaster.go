// Bloodmage Engine
// Copyright (C) 2023 Frank Mayer
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

package firstperson

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/mathf"
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	fov            = mathf.Pi / 3
	halfFov        = fov / 2
	maxDepth int32 = 20
)

var (
	numOfRays  int32
	deltaAngle float32
	scale      int32
	screenDist float32 = 0.5

	level = [][]byte{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
)

func RenderViewport() {
	ox := core.P.X
	oy := core.P.Y
	xLevel := mathf.Floor(ox)
	yLevel := mathf.Floor(oy)

	rayAngle := core.P.Angle - halfFov + mathf.Epsilon
	for ray := int32(0); ray < numOfRays; ray++ {
		sinA := mathf.Sin(rayAngle)
		cosA := mathf.Cos(rayAngle)

		var dy float32
		var dx float32
		var deltaDepth float32

		// horizontals
		var yHor float32
		if sinA > 0 {
			yHor = yLevel + 1
			dy = 1
		} else {
			yHor = yLevel - mathf.Epsilon
			dy = -1
		}
		depthHor := (yHor - oy) / sinA
		xHor := ox + depthHor*cosA
		deltaDepth = dy / sinA
		dx = deltaDepth * cosA
		var i int32
		for i = 0; i < maxDepth; i++ {
			tileX := int(mathf.Floor(xHor))
			tileY := int(mathf.Floor(yHor))
			if tileX < 0 || tileX >= len(level[0]) || tileY < 0 || tileY >= len(level) {
				break
			}
			if level[tileY][tileX] != 0 {
				break
			}
			xHor += dx
			yHor += dy
			depthHor += deltaDepth
		}

		// verticals
		var xVert float32
		if cosA > 0 {
			xVert = xLevel + 1
			dx = 1
		} else {
			xVert = xLevel - mathf.Epsilon
			dx = -1
		}
		depthVert := (xVert - ox) / cosA
		yVert := oy + depthVert*sinA
		deltaDepth = dx / cosA
		dy = deltaDepth * sinA
		for i = 0; i < maxDepth; i++ {
			tileX := int(mathf.Floor(xVert))
			tileY := int(mathf.Floor(yVert))
			if tileX < 0 || tileX >= len(level[0]) || tileY < 0 || tileY >= len(level) {
				break
			}
			if level[tileY][tileX] != 0 {
				break
			}
			xVert += dx
			yVert += dy
			depthVert += deltaDepth
		}

		// depth
		var depth float32
		if depthHor < depthVert {
			depth = depthHor
		} else {
			depth = depthVert
		}

		// remove fish eye
		depth *= mathf.Cos(core.P.Angle - rayAngle)

		// projection
		projHeight := screenDist / (depth + mathf.Epsilon)

		// draw wall
		rect := sdl.Rect{
			X: ray * scale, Y: int32(core.HalfHeightF() - projHeight/2),
			W: scale, H: int32(projHeight),
		}
		// distant walls are darker
		darkness := uint8(255 / (depth + 1))
		err := core.Renderer().SetDrawColor(darkness, darkness, darkness, 255)
		if err != nil {
			log.Error(err)
			return
		}
		err = core.Renderer().FillRect(&rect)
		if err != nil {
			log.Error(err)
			return
		}

		rayAngle += deltaAngle
	}
}
