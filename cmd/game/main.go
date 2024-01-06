// The main game
package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/charmbracelet/log"
)

func main() {
	core.InitOptions()

	l, err := level.Load("assets/levels/level1.bmlvl")
	if err != nil {
		log.Fatal(err)
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
		err = firstperson.RenderViewport()
		if err != nil {
			log.Error(err)
		}

		// draw frame
		core.Present()
	}
}
