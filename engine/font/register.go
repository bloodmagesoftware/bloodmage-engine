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

// Package font is for dealing with ttf fonts in the ui.
//
// Call font.Init() before using any other font function.
// Call font.Quit() at the end of your game.
package font

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/ttf"
)

var (
	// fonts holds all registered fonts wether they are currently loaded or not.
	fonts map[string]*font
	// defaultFont is the font that is used when no other font is specified.
	defaultFont *font
)

type font struct {
	path string
	ttf  *ttf.Font
}

// Font ensures the font is loaded and returns it.
func (f *font) Font() (*ttf.Font, error) {
	if f.ttf != nil {
		return f.ttf, nil
	}

	font, err := ttf.OpenFont(f.path, 32)
	if err != nil {
		return nil, err
	}
	f.ttf = font

	return font, nil
}

func Init() error {
	return ttf.Init()
}

// Quit closes all loaded fonts.
// After calling this function you can no longer use any font functions before calling Init() again.
func Quit() {
	for _, font := range fonts {
		if font.ttf != nil {
			font.ttf.Close()
			font.ttf = nil
		}
	}
	ttf.Quit()
}

func Register(fontPath string, name string) error {
	if fonts == nil {
		fonts = make(map[string]*font)
	}
	if _, ok := fonts[name]; ok {
		return fmt.Errorf("font %s already registered", name)
	}
	// check if file exists with os.Stat()
	if _, err := os.Stat(fontPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("font %s does not exist", fontPath)
		} else {
			return fmt.Errorf("error checking if font %s exists: %s", fontPath, err)
		}
	}
	fonts[name] = &font{
		path: fontPath,
	}
	return nil
}

func SetDefault(name string) error {
	if _, ok := fonts[name]; !ok {
		return fmt.Errorf("font %s not registered", name)
	}
	defaultFont = fonts[name]
	return nil
}

func Default() (*ttf.Font, error) {
	if defaultFont != nil {
		return defaultFont.Font()
	}
	return nil, fmt.Errorf("no default font set")
}

func Get(name string) (*ttf.Font, error) {
	if font, ok := fonts[name]; ok {
		return font.Font()
	}
	return nil, fmt.Errorf("font %s not registered", name)
}
