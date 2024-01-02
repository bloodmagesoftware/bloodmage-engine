// Bloodmage Engine
// Copyright (C) 2024 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

package textures

import (
	"image/color"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	// map of id to texture
	registry = make(map[Key]*Texture)
	// default texture for missing textures
	defaultTexture *Texture
)

func Register(texturepath string, key Key) *Texture {
	t := &Texture{
		path: texturepath,
	}

	// add texture to registry
	registry[key] = t

	return t
}

func Get(key Key) *Texture {
	t, ok := registry[key]
	if ok {
		return t
	}
	return DefaultTexture()
}

func DefaultTexture() *Texture {
	if defaultTexture != nil {
		return defaultTexture
	}

	// create pink texture for missing textures using sdl
	s, err := sdl.CreateRGBSurface(0, 1, 1, 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	s.Set(0, 0, color.RGBA{R: 255, G: 0, B: 255, A: 255})

	t, err := core.Renderer().CreateTextureFromSurface(s)
	if err != nil {
		panic(err)
	}

	s.Free()

	defaultTexture = &Texture{
		path:    "",
		width:   1,
		height:  1,
		texture: t,
	}

	return defaultTexture
}
