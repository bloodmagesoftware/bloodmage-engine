package engine

import "github.com/veandco/go-sdl2/sdl"

const (
	Forward = sdl.SCANCODE_W
	Back    = sdl.SCANCODE_S
	Left    = sdl.SCANCODE_A
	Right   = sdl.SCANCODE_D
    Escape  = sdl.SCANCODE_ESCAPE
)

func KeyDown(key uint8) bool {
    return keystates[key] != 0
}

