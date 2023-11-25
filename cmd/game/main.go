package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/level"
)

func main() {
	core.InitOptions()

	l := level.Level{
		Collision: []uint32{
			0b1111111111,
			0b1000000001,
			0b1000000001,
			0b1000000001,
			0b1000000001,
			0b1111111111,
		},
		Textures: [][]byte{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	level.Set(&l)

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
