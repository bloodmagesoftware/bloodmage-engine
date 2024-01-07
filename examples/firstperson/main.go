package main

import (
	"fmt"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/firstperson"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/font"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/textures"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/ui"
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	var err error
	core.InitOptions()

	l := level.New()
	level.Set(l)

	// register textures
	textures.Register("assets/textures/2.bmp", 2)
	textures.Register("assets/textures/1.bmp", 1)

	// register fonts
	if err = font.Init(); err != nil {
		log.Fatal(err)
	}
	if err = font.Register("./assets/fonts/GlassAntiqua-Regular.ttf", "Glass Antiqua"); err != nil {
		log.Fatal(err)
	}
	if err = font.SetDefault("Glass Antiqua"); err != nil {
		log.Fatal(err)
	}

	// set player start position
	core.P.X = 1.5
	core.P.Y = 1.5

	// inet game mode
	firstperson.Init()
	core.Start("First Person Example")
	defer core.Stop()

	core.LockCursor(false)

	err = ttf.Init()
	if err != nil {
		log.Fatal(err)
	}

	document, err := ui.Parse("./assets/ui/helloworld.xml")
	if err != nil {
		log.Fatal(err)
	}

	btnEl, ok := document.GetButtonElementById("btn")
	if !ok {
		log.Fatal("Could not find element with id 'btn'")
	}
	i := 0

	counterEl, ok := document.GetTextElementById("counter")
	if !ok {
		log.Fatal("Could not find element with id 'counter'")
	}

	// game loop
	for core.Running() {
		if core.KeyStates()[sdl.SCANCODE_ESCAPE] != 0 {
			break
		}
		firstperson.MovePlayer()
		if err = firstperson.RenderViewport(); err != nil {
			log.Fatal(err)
		}

		if err = document.Draw(); err != nil {
			log.Fatal(err)
		}

		if btnEl.Clicked() {
			i++
			_ = counterEl.SetContent(fmt.Sprintf("Clicked %d times", i))
		}

		// draw frame
		core.Present()
	}
}
