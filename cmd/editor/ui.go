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

package main

import (
	"errors"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/font"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/textures"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/ui"
)

var (
	// UI is the main UI document
	UI *ui.Document
	// FileEl is the file path element
	FileEl *ui.Text
	// SaveEl is the save button element
	SaveEl *ui.Button
	// ExitEl is the exit button element
	ExitEl *ui.Button
)

func InitUi() error {
	// register textures
	textures.Register("assets/textures/1.bmp", 1)
	textures.Register("assets/textures/2.bmp", 2)

	// register fonts
	if err := font.Init(); err != nil {
		return err
	}
	if err := font.Register("./assets/fonts/Roboto-Regular.ttf", "Roboto"); err != nil {
		return err
	}
	if err := font.SetDefault("Roboto"); err != nil {
		return err
	}

	doc, err := ui.Parse("./cmd/editor/ui.xml")
	if err != nil {
		return err
	}

	UI = doc

	var ok bool
	if FileEl, ok = doc.GetTextElementById("file"); !ok {
		return errors.New("failed to get file element from ui")
	}
	if SaveEl, ok = doc.GetButtonElementById("save"); !ok {
		return errors.New("failed to get save element from ui")
	}
	if ExitEl, ok = doc.GetButtonElementById("exit"); !ok {
		return errors.New("failed to get exit element from ui")
	}

	return nil
}

func UpdateUi() error {
	if err := UI.Draw(); err != nil {
		return err
	}

	txt := "File: " + levelFile
	if unsavedChanges {
		txt += "*"
	}
	if err := FileEl.SetContent(txt); err != nil {
		return err
	}

	return nil
}
