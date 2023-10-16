package main

import (
	"game.frankmayer.io/engine"
)

func main() {
	engine.Start("Bloodmage Engine")
	defer engine.Quit()

	for engine.Running() {
        engine.MovePlayer()
		engine.RenderViewport()
		if engine.KeyDown(engine.KeyEscape) {
			break
		}
		engine.Present()
	}
}
