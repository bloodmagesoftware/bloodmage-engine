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

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	doc       *document
	id        string
	mouseDown bool
	content   Element
	rect      sdl.Rect
}

func newButton() *Button {
	return &Button{
		doc:       nil,
		id:        "",
		mouseDown: false,
		content:   nil,
		rect:      sdl.Rect{},
	}
}

func (b *Button) AppendChild(e Element) error {
	if b.content != nil {
		return errors.New("button already has content")
	}
	b.content = e
	return nil
}

func (b *Button) SetAttribute(key, value string) error {
	switch key {
	case "id":
		b.id = value
	default:
		return errors.New("unknown attribute: " + key)
	}
	return nil
}

func (b *Button) setDocument(doc *document) {
	b.doc = doc
}

func (b *Button) setTextContent(content string) error {
	return errors.New("button cannot have text content")
}

func (b *Button) setRect(rect *sdl.Rect) {
	b.rect.X = rect.X
	b.rect.Y = rect.Y
	b.rect.W = rect.W
	b.rect.H = rect.H
}

func (b *Button) MouseOver() bool {
	if b.rect.X <= core.MouseX && core.MouseX <= b.rect.X+b.rect.W && b.rect.Y <= core.MouseY && core.MouseY <= b.rect.Y+b.rect.H {
		core.NotifyCursorHover()
		return true
	}
	b.mouseDown = false
	return false
}

// Clicked returns true if these is a rising edge of the left mouse button
func (b *Button) Clicked() bool {
	lMouseDown := core.MouseState&sdl.ButtonLMask() != 0
	if b.MouseOver() {
		if lMouseDown {
			b.mouseDown = true
			return false
		} else {
			if b.mouseDown {
				b.mouseDown = false
				return true
			} else {
				return false
			}
		}
	} else {
		// b.mouseDown = false
		return false
	}
}
