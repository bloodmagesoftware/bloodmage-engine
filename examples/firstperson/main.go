package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/level"
)

func main() {
	core.InitOptions()

	l := level.New()
	level.Set(l)

	core.P.X = 1.5
	core.P.Y = 1.5

	firstperson.Init()
	core.Start("First Person Example")
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
