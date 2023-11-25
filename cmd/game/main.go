package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/firstperson"
)

func main() {
	core.InitOptions()
	firstperson.Init()
	core.Start("Bloodmage Engine")
	defer core.Stop()

	core.LockCursor(true)

	// game loop
	for core.Running() {
		firstperson.GetMouseInput()
		firstperson.MovePlayer()
		firstperson.RenderViewport()

		// draw frame
		core.Present()
	}
}
