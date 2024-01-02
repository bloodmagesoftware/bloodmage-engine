package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/textures"
	"github.com/charmbracelet/log"
)

func main() {
	core.InitOptions()

	l := level.New()
	level.Set(l)
	textures.Register("assets/textures/0.bmp", 0)
	textures.Register("assets/textures/1.bmp", 1)

	core.P.X = 1.5
	core.P.Y = 1.5

	firstperson.Init()
	core.Start("First Person Example")
	defer core.Stop()

	var err error

	core.LockCursor(true)

	// game loop
	for core.Running() {
		firstperson.GetMouseInput()
		firstperson.MovePlayer()
		err = firstperson.RenderViewport()
		if err != nil {
			log.Error(err)
		}

		// draw frame
		core.Present()
	}
}
