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

package firstperson

import (
	"errors"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/constants"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/textures"
	"github.com/chewxy/math32"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	fov            = math32.Pi / 3
	halfFov        = fov / 2
	maxDepth int32 = 20
)

var (
	numOfRays  int32
	deltaAngle float32
	scale      int32
	screenDist float32 = 0.5
)

func RenderViewport() error {
	var err error
	err = renderCeiling()
	if err != nil {
		return errors.Join(
			errors.New("failed to render ceiling"),
			err,
		)
	}
	err = renderFloor()
	if err != nil {
		return errors.Join(
			errors.New("failed to render floor"),
			err,
		)
	}
	err = renderWalls()
	if err != nil {
		return errors.Join(
			errors.New("failed to render walls"),
			err,
		)
	}
	return nil
}

func renderFloor() error {
	renderColor := level.FloorTexture()

	dstRect := sdl.Rect{
		X: 0, Y: int32(core.HalfHeightF()),
		W: int32(core.Width()), H: int32(core.HalfHeightF()),
	}

	err := core.Renderer().Copy(renderColor, nil, &dstRect)
	return err
}

func renderCeiling() error {
	renderColor := level.CeilingTexture()

	dstRect := sdl.Rect{
		X: 0, Y: 0,
		W: int32(core.Width()), H: int32(core.HalfHeightF()),
	}

	err := core.Renderer().Copy(renderColor, nil, &dstRect)
	return err
}

func renderWalls() error {
	var err error
	var renderColor *sdl.Texture

	ox := core.P.X
	oy := core.P.Y
	xLevel := math32.Floor(ox)
	yLevel := math32.Floor(oy)

	rayAngle := core.P.Angle - halfFov + constants.Epsilon
	for ray := int32(0); ray < numOfRays; ray++ {
		sinA := math32.Sin(rayAngle)
		cosA := math32.Cos(rayAngle)
		var textureHor *textures.Texture
		var textureVert *textures.Texture

		var dy float32
		var dx float32
		var deltaDepth float32

		// horizontals
		var yHor float32
		if sinA > 0 {
			yHor = yLevel + 1
			dy = 1
		} else {
			yHor = yLevel - constants.Epsilon
			dy = -1
		}
		depthHor := (yHor - oy) / sinA
		xHor := ox + depthHor*cosA
		deltaDepth = dy / sinA
		dx = deltaDepth * cosA
		var i int32
		for i = 0; i < maxDepth; i++ {
			tileX := int(math32.Floor(xHor))
			tileY := int(math32.Floor(yHor))
			if level.Collision(tileX, tileY) {
				textureHor = level.WallTexture(tileX, tileY)
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
			xVert = xLevel - constants.Epsilon
			dx = -1
		}
		depthVert := (xVert - ox) / cosA
		yVert := oy + depthVert*sinA
		deltaDepth = dx / cosA
		dy = deltaDepth * sinA
		for i = 0; i < maxDepth; i++ {
			tileX := int(math32.Floor(xVert))
			tileY := int(math32.Floor(yVert))
			if level.Collision(tileX, tileY) {
				textureVert = level.WallTexture(tileX, tileY)
				break
			}
			xVert += dx
			yVert += dy
			depthVert += deltaDepth
		}

		// depth
		var depth float32
		var texture *textures.Texture
		var offset float32
		if depthHor < depthVert {
			depth = depthHor
			texture = textureHor
			xHor = math32.Mod(xHor, 1)
			if sinA > 0 {
				offset = 1 - xHor
			} else {
				offset = xHor
			}
		} else {
			depth = depthVert
			texture = textureVert
			yVert = math32.Mod(yVert, 1)
			if cosA > 0 {
				offset = yVert
			} else {
				offset = 1 - yVert
			}
		}

		// remove fish eye
		depth *= math32.Cos(core.P.Angle - rayAngle)

		// projection
		projHeight := screenDist / (depth + constants.Epsilon)

		// draw wall
		dstRect := sdl.Rect{
			X: ray * scale, Y: int32(core.HalfHeightF() - projHeight/2),
			W: scale, H: int32(projHeight),
		}
		srcRect := sdl.Rect{
			X: int32(offset * float32(texture.Width())),
			Y: 0,
			W: 1,
			H: int32(texture.Height()),
		}
		renderColor, err = texture.Texture()
		if err != nil {
			return err
		}
		err = core.Renderer().Copy(renderColor, &srcRect, &dstRect)
		if err != nil {
			return err
		}
		// distant walls are darker
		darkness := uint8(int32((depth+1)*255) / maxDepth)
		err = core.Renderer().SetDrawColor(0, 0, 0, darkness)
		if err != nil {
			return err
		}
		err = core.Renderer().FillRect(&dstRect)
		if err != nil {
			return err
		}

		rayAngle += deltaAngle
	}

	return nil
}
