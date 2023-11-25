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
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/level"
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
		core.P.Speed = speed
	} else if KeyDown(KeyBack) {
		core.P.Speed = -speed
	} else {
		core.P.Speed = 0
	}

	if KeyDown(KeyLeft) {
		core.P.Strafe = -speed
	} else if KeyDown(KeyRight) {
		core.P.Strafe = speed
	} else {
		core.P.Strafe = 0
	}

	var x, y int
	xDir := core.P.Speed*math32.Cos(core.P.Angle) + core.P.Strafe*math32.Cos(core.P.Angle+math.Pi/2)
	newX := core.P.X + xDir*core.DeltaTime
	bufferX := newX + xDir*0.01
	y = int(math32.Floor(core.P.Y))
	x = int(math32.Floor(bufferX))
	if !level.Collision(x, y) {
		core.P.X = newX
	}

	yDir := core.P.Speed*math32.Sin(core.P.Angle) + core.P.Strafe*math32.Sin(core.P.Angle+math32.Pi/2)
	newY := core.P.Y + yDir*core.DeltaTime
	bufferY := newY + yDir*0.01
	x = int(math32.Floor(core.P.X))
	y = int(math32.Floor(bufferY))
	if !level.Collision(x, y) {
		core.P.Y = newY
	}
}
