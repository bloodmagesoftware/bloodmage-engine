package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine"
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
