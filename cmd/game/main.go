package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/level"
)

func main() {
	core.InitOptions()

	l, err := level.Load("assets/levels/level1.bmlvl")
	if err != nil {
		panic(err)
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
