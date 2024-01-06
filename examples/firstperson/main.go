package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/textures"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/ui"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	core.InitOptions()

	l := level.New()
	level.Set(l)
	textures.Register("assets/textures/2.bmp", 2)
	textures.Register("assets/textures/1.bmp", 1)

	core.P.X = 1.5
	core.P.Y = 1.5

	firstperson.Init()
	core.Start("First Person Example")
	defer core.Stop()

	var err error

	core.LockCursor(true)

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	font, err := ttf.OpenFont("assets/fonts/GlassAntiqua-Regular.ttf", 16)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	document, err := ui.Parse("./assets/ui/helloworld.xml")
	if err != nil {
		panic(err)
	}

	// game loop
	for core.Running() {
		if core.KeyStates()[sdl.SCANCODE_ESCAPE] != 0 {
			break
		}
		firstperson.GetMouseInput()
		firstperson.MovePlayer()
		if err = firstperson.RenderViewport(); err != nil {
			panic(err)
		}

		surface, err := font.RenderUTF8Solid("Hällo Wörld!", sdl.Color{R: 255, G: 255, B: 255, A: 255})
		if err != nil {
			panic(err)
		}
		defer surface.Free()

		texture, err := core.Renderer().CreateTextureFromSurface(surface)
		if err != nil {
			panic(err)
		}

		defer texture.Destroy()

		if err = core.Renderer().Copy(texture, nil, &sdl.Rect{X: 0, Y: 0, W: 600, H: 100}); err != nil {
			panic(err)
		}

		if err = ui.Draw(core.Renderer(), document); err != nil {
			panic(err)
		}

		// draw frame
		core.Present()
	}
}
