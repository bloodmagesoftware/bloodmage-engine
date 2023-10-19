package main

import (
	"game.frankmayer.io/engine"
)

func main() {
	engine.InitOptions("bloodmage-engine")
	engine.Start("Bloodmage Engine")
	defer engine.Stop()

	engine.SetCursorLock(true)

	for engine.Running() {
		engine.MovePlayer()
		engine.RenderViewport()

		if engine.IsCursorLocked() {
			if engine.KeyDown(engine.KeyEscape) {
				engine.SetCursorLock(false)
			}
		} else {
			test_text := engine.CreateAlignedText(
				"Continue",
				0.5, 0.5,
				engine.UI_ALIGN_CENTER, engine.UI_ALIGN_CENTER,
			)
			test_text.Draw()
			if test_text.MouseDown() {
				engine.SetCursorLock(true)
			}
		}
		engine.Present()
	}
}
