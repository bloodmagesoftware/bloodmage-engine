package main

import (
	"game.frankmayer.io/engine"
)

func main() {
	engine.InitOptions("bloodmage-engine")
	engine.Start("Bloodmage Engine")
	defer engine.Stop()

	engine.SetCursorLock(true)

	for engine.Running() {
		engine.MovePlayer()
		engine.RenderViewport()
		if engine.KeyDown(engine.KeyEscape) {
			engine.SetCursorLock(false)
		}
		engine.Present()
	}
}
