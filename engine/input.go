package engine

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	KeyForward   = sdl.SCANCODE_W
	KeyBack      = sdl.SCANCODE_S
	KeyLeft      = sdl.SCANCODE_A
	KeyRight     = sdl.SCANCODE_D
	KeyTurnRight = sdl.SCANCODE_E
	KeyTurnLeft  = sdl.SCANCODE_Q
	KeyEscape    = sdl.SCANCODE_ESCAPE
	speed        = 1
)

func KeyDown(key uint8) bool {
	return keystates[key] != 0
}

func applyInputEvents() {
	if KeyDown(KeyTurnRight) {
		P.Angle += 0.1
	}
	if KeyDown(KeyTurnLeft) {
		P.Angle -= 0.1
	}

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

	if KeyDown(KeyTurnRight) {
		P.Angle += 0.1
	} else if KeyDown(KeyTurnLeft) {
		P.Angle -= 0.1
	} else {
		P.Angle = P.Angle
	}
}

func MovePlayer() {
	applyInputEvents()
	P.X += P.Speed * math.Cos(P.Angle)
	P.Y += P.Speed * math.Sin(P.Angle)
	P.X += P.Strafe * math.Cos(P.Angle+math.Pi/2)
	P.Y += P.Strafe * math.Sin(P.Angle+math.Pi/2)
}
