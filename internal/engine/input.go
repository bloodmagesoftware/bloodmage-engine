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
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	KeyForward         = sdl.SCANCODE_W
	KeyBack            = sdl.SCANCODE_S
	KeyLeft            = sdl.SCANCODE_A
	KeyRight           = sdl.SCANCODE_D
	KeyEscape          = sdl.SCANCODE_ESCAPE
	speed      float64 = 3
	turnSpeed  float64 = 0.1
)

var (
	MouseY      int32
	MouseX      int32
	MouseDeltaY int32
	MouseDeltaX int32
	MouseState  uint32
)

func getMouseInput() {
	MouseX, MouseY, MouseState = sdl.GetMouseState()
	if cursorLocked {
		MouseDeltaX = MouseX - centerX
		MouseDeltaY = MouseY - centerY
		window.WarpMouseInWindow(centerX, centerY)
	} else {
		MouseDeltaX = 0
		MouseDeltaY = 0
	}
}

func KeyDown(key uint8) bool {
	return keyStates[key] != 0
}

func MovePlayer() {
	P.Angle += float64(MouseDeltaX) * turnSpeed * DeltaTime

	if KeyDown(KeyForward) {
		P.Speed = speed
	} else if KeyDown(KeyBack) {
		P.Speed = -speed
	} else {
		P.Speed = 0
	}

	if KeyDown(KeyLeft) {
		P.Strafe = -speed
	} else if KeyDown(KeyRight) {
		P.Strafe = speed
	} else {
		P.Strafe = 0
	}

	var x, y int
	xDir := P.Speed*math.Cos(P.Angle) + P.Strafe*math.Cos(P.Angle+math.Pi/2)
	newX := P.X + xDir*DeltaTime
	bufferX := newX + xDir*0.01
	y = int(math.Floor(P.Y))
	x = int(math.Floor(bufferX))
	if x > 0 && x < len(level[y]) && level[y][x] == 0 {
		P.X = newX
	}

	yDir := P.Speed*math.Sin(P.Angle) + P.Strafe*math.Sin(P.Angle+math.Pi/2)
	newY := P.Y + yDir*DeltaTime
	bufferY := newY + yDir*0.01
	x = int(math.Floor(P.X))
	y = int(math.Floor(bufferY))
	if y > 0 && y < len(level) && level[y][x] == 0 {
		P.Y = newY
	}
}

func SetCursorLock(lock bool) {
	cursorLocked = lock
	if cursorLocked {
		_, err := sdl.ShowCursor(sdl.DISABLE)
		if err != nil {
			log.Error(err)
			return
		}
		sdl.SetRelativeMouseMode(true)
		window.SetGrab(true)
		window.WarpMouseInWindow(centerX, centerY)
	} else {
		_, err := sdl.ShowCursor(sdl.ENABLE)
		if err != nil {
			log.Error(err)
			return
		}
		sdl.SetRelativeMouseMode(false)
		window.SetGrab(false)
	}
}

func IsCursorLocked() bool {
	return cursorLocked
}
