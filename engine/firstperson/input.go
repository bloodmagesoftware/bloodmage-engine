// Bloodmage Engine - Retro first person game engine
// Copyright (C) 2024  Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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

func KeyDown(key uint8) bool {
	return core.KeyStates()[key] != 0
}

func MovePlayer() {
	core.P.Angle += float32(core.MouseDeltaX) * turnSpeed * core.DeltaTime

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
