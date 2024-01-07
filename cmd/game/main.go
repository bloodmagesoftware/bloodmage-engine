// Bloodmage Engine - Retro first person game engine
// Copyright (C) 2024  Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// The main game
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
)

func main() {
	var err error
	core.InitOptions()

	l, err := level.Load("assets/levels/demo")
	if err != nil {
		log.Fatal(err)
	}
	level.Set(l)

	// register textures
	textures.Register("assets/textures/1.bmp", 1)
	textures.Register("assets/textures/2.bmp", 2)

	// register fonts
	if err = font.Init(); err != nil {
		log.Fatal(err)
	}
	if err = font.Register("./assets/fonts/Roboto-Regular.ttf", "Roboto"); err != nil {
		log.Fatal(err)
	}
	if err = font.SetDefault("Roboto"); err != nil {
		log.Fatal(err)
	}

	// set player start position
	core.P.X = 1.5
	core.P.Y = 1.5

	// init game mode
	firstperson.Init()
	core.Start("First Person Example")
	defer core.Stop()

	core.LockCursor(true)

	document, err := ui.Parse("./assets/ui/helloworld.xml")
	if err != nil {
		log.Fatal(err)
	}

	btnEl, ok := document.GetButtonElementById("clickme")
	if !ok {
		log.Fatal("Could not find element with id 'btn'")
	}

	exitBtnEl, ok := document.GetButtonElementById("exit")
	if !ok {
		log.Fatal("Could not find element with id 'exit'")
	}

	i := 0

	counterEl, ok := document.GetTextElementById("counter")
	if !ok {
		log.Fatal("Could not find element with id 'counter'")
	}

	escDown := false

	// game loop
	for core.Running() {
		if core.KeyStates()[sdl.SCANCODE_ESCAPE] != 0 {
			if !escDown {
				core.LockCursor(!core.IsCursorLocked())
			}
			escDown = true
		} else {
			escDown = false
		}

		if err = firstperson.RenderViewport(); err != nil {
			log.Fatal(err)
		}

		if core.IsCursorLocked() {
			firstperson.MovePlayer()
		} else {
			if err = document.Draw(); err != nil {
				log.Fatal(err)
			}
			if btnEl.Clicked() {
				i++
				_ = counterEl.SetContent(fmt.Sprintf("Clicked %d times", i))
			}
			if exitBtnEl.Clicked() {
				break
			}
		}

		// draw frame
		core.Present()
	}
}
