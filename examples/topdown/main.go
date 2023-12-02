package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/level"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/topdown"
)

func main() {
	core.InitOptions()

	l := level.New()
	level.Set(l)

	core.P.X = 1.5
	core.P.Y = 1.5

	topdown.Init()
	core.Start("Top Down Example")
	defer core.Stop()

	core.LockCursor(true)

	// game loop
	for core.Running() {
		topdown.ProcessInput()
		topdown.MovePlayer()
		topdown.RenderViewport()

		// draw frame
		core.Present()
	}
}
