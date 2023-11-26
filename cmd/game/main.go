package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/level"
	"log"
)

func main() {
	core.InitOptions()

	l, err := level.Load("assets/levels/level1.pb.bin")
	if err != nil {
		log.Panicln("Failed to load level", "error", err)
	}
	level.Set(l)

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
