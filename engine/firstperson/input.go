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
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/chewxy/math32"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	KeyForward         = sdl.SCANCODE_W
	KeyBack            = sdl.SCANCODE_S
	KeyLeft            = sdl.SCANCODE_A
	KeyRight           = sdl.SCANCODE_D
	speed      float32 = 3
	turnSpeed  float32 = 0.1
)

var (
	MouseX      int32
	MouseDeltaX int32
	MouseState  uint32
)

func GetMouseInput() {
	MouseX, _, MouseState = sdl.GetMouseState()
	if core.IsCursorLocked() {
		MouseDeltaX = MouseX - core.CenterX()
		core.Window().WarpMouseInWindow(core.CenterX(), core.CenterY())
	} else {
		MouseDeltaX = 0
	}
}

func KeyDown(key uint8) bool {
	return core.KeyStates()[key] != 0
}

func MovePlayer() {
	core.P.Angle += float32(MouseDeltaX) * turnSpeed * core.DeltaTime

	if KeyDown(KeyForward) {
		core.P.Speed = 1
	} else if KeyDown(KeyBack) {
		core.P.Speed = -1
	} else {
		core.P.Speed = 0
	}

	if KeyDown(KeyLeft) {
		core.P.Strafe = -1
	} else if KeyDown(KeyRight) {
		core.P.Strafe = 1
	} else {
		core.P.Strafe = 0
	}

	xDir := core.P.Speed*math32.Cos(core.P.Angle) + core.P.Strafe*math32.Cos(core.P.Angle+math.Pi/2)
	yDir := core.P.Speed*math32.Sin(core.P.Angle) + core.P.Strafe*math32.Sin(core.P.Angle+math32.Pi/2)
	length := math32.Sqrt(xDir*xDir + yDir*yDir)
	if length == 0 {
		return
	}
	xDir /= length
	yDir /= length

	newX := core.P.X + xDir*core.DeltaTime*speed
	newY := core.P.Y + yDir*core.DeltaTime*speed

	if !level.Collision(int(newX), int(core.P.Y)) {
		core.P.X = newX
	}
	if !level.Collision(int(core.P.X), int(newY)) {
		core.P.Y = newY
	}
}
