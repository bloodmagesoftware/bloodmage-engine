package main

import (
	"game.frankmayer.io/engine"
)

func main() {
	engine.Start("Bloodmage Engine", false)
	defer engine.Stop()

	for engine.Running() {
		engine.MovePlayer()
		engine.RenderViewport()
		if engine.KeyDown(engine.KeyEscape) {
			break
		}
		engine.Present()
	}
}
