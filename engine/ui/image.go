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

package ui

import (
	"errors"
	"strconv"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	doc     *document
	id      string
	src     string
	width   int32
	height  int32
	texture *sdl.Texture
}

func newImage() *Image {
	return &Image{
		doc:     nil,
		id:      "",
		src:     "",
		width:   -1,
		height:  -1,
		texture: nil,
	}
}

func (i *Image) AppendChild(e Element) error {
	return errors.New("Image cannot have children")
}

func (i *Image) SetAttribute(key, value string) error {
	switch key {
	case "id":
		i.id = value
	case "src":
		return i.SetSrc(value)
	case "width":
		if w, err := strconv.Atoi(value); err != nil {
			return err
		} else if w < 0 {
			return errors.New("width must be greater than 0")
		} else {
			i.width = int32(w)
		}
	case "height":
		if h, err := strconv.Atoi(value); err != nil {
			return err
		} else if h < 0 {
			return errors.New("height must be greater than 0")
		} else {
			i.height = int32(h)
		}
	default:
		return errors.New("unknown attribute: " + key)
	}
	return nil
}

func (i *Image) setDocument(doc *document) {
	i.doc = doc
}

func (i *Image) Src() string {
	return i.src
}

func (i *Image) SetSrc(src string) error {
	if i.src == src {
		return nil
	}
	if i.texture != nil {
		err := i.texture.Destroy()
		if err != nil {
			return err
		}
	}
	surface, err := sdl.LoadBMP(src)
	if err != nil {
		return err
	}
	i.texture, err = core.Renderer().CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}
	surface.Free()
	i.src = src
	return nil
}

func (i *Image) Texture() (*sdl.Texture, error) {
	if i.texture == nil {
		return nil, errors.New("texture is nil")
	}
	return i.texture, nil
}

func (i *Image) rect() (*sdl.Rect, error) {
	if i.width <= 0 {
		return nil, errors.New("image width must be greater than 0")
	}
	if i.height <= 0 {
		return nil, errors.New("image height must be greater than 0")
	}
	return &sdl.Rect{
		X: 0, Y: 0,
		W: i.width, H: i.height,
	}, nil
}

func (i *Image) setTextContent(content string) error {
	return errors.New("Image cannot have text content")
}
