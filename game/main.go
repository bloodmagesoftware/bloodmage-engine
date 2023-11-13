package main

import (
	"bloodmagesoftware.io/engine"
)

func main() {
	engine.InitOptions()
	engine.Start("Bloodmage Engine")
	defer engine.Stop()

	engine.SetCursorLock(true)

	// game loop
	for engine.Running() {
		engine.MovePlayer()
		engine.RenderViewport()

		// draw frame
		engine.Present()
	}
}
