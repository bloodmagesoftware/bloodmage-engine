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
	"github.com/chewxy/math32"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	velocity = 2
)

func ProcessInput() {
	core.P.Speed = 0.0
	core.P.Strafe = 0.0

	if core.KeyStates()[sdl.SCANCODE_DOWN] == 1 {
		core.P.Speed = 1
	} else if core.KeyStates()[sdl.SCANCODE_UP] == 1 {
		core.P.Speed = -1
	}

	if core.KeyStates()[sdl.SCANCODE_LEFT] == 1 {
		core.P.Strafe = -1
	} else if core.KeyStates()[sdl.SCANCODE_RIGHT] == 1 {
		core.P.Strafe = 1
	}
}

func MovePlayer() {
	// calculate direction vector (normalized)
	length := math32.Sqrt(core.P.Speed*core.P.Speed + core.P.Strafe*core.P.Strafe)
	if length == 0 {
		return
	}

	dirX := core.P.Strafe / length * velocity
	dirY := core.P.Speed / length * velocity

	// calculate new position
	newX := core.P.X + dirX*core.DeltaTime
	newY := core.P.Y + dirY*core.DeltaTime

	// check if new position is valid
	if !level.CollisionF(newX, core.P.Y) {
		core.P.X = newX
	}

	if !level.CollisionF(core.P.X, newY) {
		core.P.Y = newY
	}
}
