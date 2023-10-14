package main

import (
	"game.frankmayer.io/engine"
)

func main() {
    engine.Start("Bloodmage Engine")
    defer engine.Quit()

    for engine.Running() {
        engine.RenderViewport()
        if engine.KeyDown(engine.Escape) {
            break
        }
        engine.Present()
    }
}
