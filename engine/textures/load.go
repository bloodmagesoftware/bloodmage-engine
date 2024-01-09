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

package textures

import (
	"errors"
	"sync"
	"time"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

func (t *Texture) load() error {
	loaderMutex.Lock()
	defer loaderMutex.Unlock()

	t.lastUsed = time.Now()

	if t.texture != nil {
		return nil
	}

	surface, err := sdl.LoadBMP(t.path)
	if err != nil {
		return err
	}
	defer surface.Free()
	t.texture, err = core.Renderer().CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}
	t.height = surface.H
	t.width = surface.W

	return nil
}

func (t *Texture) unload() error {
	if t.texture != nil {
		err := t.texture.Destroy()
		if err != nil {
			return err
		}
		t.texture = nil
	}

	return nil
}

func (t *Texture) Texture() (*sdl.Texture, error) {
	err := t.load()
	return t.texture, err
}

var loaderMutex = sync.Mutex{}

const oldTextureTimeout = 5 * time.Second

func monitor() {
	for {
		time.Sleep(10 * time.Second)
		gc()
	}
}

func gc() {
	loaderMutex.Lock()
	defer loaderMutex.Unlock()
	now := time.Now()
	for _, t := range registry {
		if t.texture == nil {
			continue
		}
		age := now.Sub(t.lastUsed)
		log.Debug("Checking texture for unload", "path", t.path, "age", age)
		if age > oldTextureTimeout {
			log.Debug("Queueing texture for unload", "path", t.path)
			t := t
			core.RunOnMainThread(func() error {
				if t == nil || t.texture == nil {
					return nil
				}
				err := t.unload()
				if err != nil {
					return errors.Join(
						errors.New("Error destroying texture"),
						err,
					)
				} else {
					log.Debug("Unloaded texture", "path", t.path)
				}
				return nil
			})
		}
	}

}
