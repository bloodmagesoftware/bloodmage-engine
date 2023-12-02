package topdown

import (
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/level"
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
