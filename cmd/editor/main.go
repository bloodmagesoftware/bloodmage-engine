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

// level editor
package main

import (
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/charmbracelet/log"
)

func main() {
	if err := InitLevel(); err != nil {
		log.Fatal(err)
	}

	if err := InitUi(); err != nil {
		log.Fatal(err)
	}

	// start editor
	core.InitOptions()
	core.Options().Fullscreen = false
	core.Start("Bloodmage Engine - Editor :: " + levelFile)
	defer core.Stop()

	// game loop
	for core.Running() {
		if err := DrawLevel(); err != nil {
			log.Fatal(err)
		}

		if err := UpdateUi(); err != nil {
			log.Fatal(err)
		}

		if ExitEl.Clicked() {
			break
		}

		if SaveEl.Clicked() {
			SaveLevel()
		}

		if !ExitEl.MouseOver() && !SaveEl.MouseOver() {
			EditLevel()
		}

		// draw frame
		core.Present()
	}
}
