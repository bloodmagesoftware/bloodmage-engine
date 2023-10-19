package main

import (
	"fmt"
	"math"

	"bloodmagesoftware.io/engine"
)

func main() {
	engine.InitOptions("bloodmage-engine")
	engine.Start("Bloodmage Engine")
	defer engine.Stop()

	engine.SetCursorLock(true)

	// game loop
	for engine.Running() {
		engine.MovePlayer()
		engine.RenderViewport()

		// is game focused?
		if engine.IsCursorLocked() {
			// is escape key pressed?
			if engine.KeyDown(engine.KeyEscape) {
				// unfocus game
				engine.SetCursorLock(false)
			}
			// fps counter
			time_el := engine.CreateAlignedText(
				fmt.Sprintf("%d FPS", int(math.Round(1/engine.DeltaTime))),
				0, 0,
				engine.UI_ALIGN_START, engine.UI_ALIGN_START,
			)
			time_el.Draw()
		} else {
			// create pause menu
			menu_elems := engine.CreateOptions("Resume", "Quit")
			// draw all menu elements
			for _, e := range menu_elems {
				e.Draw()
			}
			if menu_elems[0].MouseDown() {
				// if "Resume" is clicked, focus game
				engine.SetCursorLock(true)
			} else if menu_elems[1].MouseDown() {
				// if "Quit" is clicked, exit game loop
				break
			}
		}

		// draw frame
		engine.Present()
	}
}
