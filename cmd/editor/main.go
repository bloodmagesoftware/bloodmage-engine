package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/firstperson"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	core.InitOptions()
	core.Options().Fullscreen = false

	firstperson.Init()
	core.Start("Bloodmage Editor")
	defer core.Stop()

	core.LockCursor(false)

	// game loop
	for core.Running() {
		if core.IsCursorLocked() && core.KeyStates()[sdl.SCANCODE_ESCAPE] != 0 {
			core.LockCursor(false)
		} else if x, y, s := sdl.GetMouseState(); (x&y) != 0 && s&sdl.MOUSEBUTTONDOWN != 0 {
			core.LockCursor(true)
		}
		firstperson.GetMouseInput()
		firstperson.MovePlayer()
		firstperson.RenderViewport()

		// draw frame
		core.Present()
	}
}
