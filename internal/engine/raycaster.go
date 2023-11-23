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

package engine

import (
	"github.com/charmbracelet/log"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	fov      float64 = math.Pi / 3
	halfFov          = fov / 2
	maxDepth         = 20
	epsilon          = 1e-6
)

var (
	numOfRays  int32
	deltaAngle float64
	scale      int32

	level = [][]int{
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
	ox := P.X
	oy := P.Y
	xLevel := math.Floor(ox)
	yLevel := math.Floor(oy)

	rayAngle := P.Angle - halfFov + epsilon
	for ray := int32(0); ray < numOfRays; ray++ {
		sinA := math.Sin(rayAngle)
		cosA := math.Cos(rayAngle)

		var dy float64
		var dx float64
		var deltaDepth float64

		// horizontals
		var yHor float64
		if sinA > 0 {
			yHor = yLevel + 1
			dy = 1.0
		} else {
			yHor = yLevel - epsilon
			dy = -1.0
		}
		depthHor := (yHor - oy) / sinA
		xHor := ox + depthHor*cosA
		deltaDepth = dy / sinA
		dx = deltaDepth * cosA
		for i := 0; i < maxDepth; i++ {
			tileX := int(math.Floor(xHor))
			tileY := int(math.Floor(yHor))
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
		var xVert float64
		if cosA > 0 {
			xVert = xLevel + 1
			dx = 1.0
		} else {
			xVert = xLevel - epsilon
			dx = -1.0
		}
		depthVert := (xVert - ox) / cosA
		yVert := oy + depthVert*sinA
		deltaDepth = dx / cosA
		dy = deltaDepth * sinA
		for i := 0; i < maxDepth; i++ {
			tileX := int(math.Floor(xVert))
			tileY := int(math.Floor(yVert))
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
		var depth float64
		if depthHor < depthVert {
			depth = depthHor
		} else {
			depth = depthVert
		}

		// remove fish eye
		depth *= math.Cos(P.Angle - rayAngle)

		// projection
		projHeight := screenDist / (depth + epsilon)

		// draw wall
		rect := sdl.Rect{
			X: ray * scale, Y: int32(halfHeightF64 - projHeight/2),
			W: scale, H: int32(projHeight),
		}
		// distant walls are darker
		darkness := uint8(255 / (depth + 1))
		err := renderer.SetDrawColor(darkness, darkness, darkness, 255)
		if err != nil {
			log.Error(err)
			return
		}
		err = renderer.FillRect(&rect)
		if err != nil {
			log.Error(err)
			return
		}

		rayAngle += deltaAngle
	}
}
