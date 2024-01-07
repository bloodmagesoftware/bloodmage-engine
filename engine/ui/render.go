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

package ui

import (
	"errors"
	"fmt"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/font"
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type drawFn func(dest *sdl.Rect) error

func (document *Document) Draw() error {
	fn, rect, err := document.root.draw()
	if err != nil {
		return err
	}
	if rect == nil {
		rect = core.ScreenRect()
	}
	if err := fn(rect); err != nil {
		return err
	}
	return nil
}

// draw functions

func (image *Image) draw() (drawFn, *sdl.Rect, error) {
	t, err := image.Texture()
	if err != nil {
		return nil, defaultRect, err
	}
	rect, err := image.rect()
	if err != nil {
		return nil, defaultRect, err
	}

	return func(dest *sdl.Rect) error {
		return core.Renderer().Copy(t, nil, dest)
	}, rect, nil
}

func (button *Button) draw() (drawFn, *sdl.Rect, error) {
	if button.content == nil {
		return nil, defaultRect, errors.New("Button content is nil")
	}

	fn, rect, err := button.content.draw()
	return func(dest *sdl.Rect) error {
		button.setRect(dest)
		return fn(dest)
	}, rect, err
}

func (list *List) draw() (drawFn, *sdl.Rect, error) {
	fnList := make([]drawFn, len(list.items))
	rectList := make([]*sdl.Rect, len(list.items))

	expectedSize := sdl.Rect{
		X: 0, Y: 0,
		W: 0, H: 0,
	}

	for i, child := range list.items {
		fn, srcRect, err := child.draw()
		if err != nil {
			return nil, defaultRect, errors.Join(
				fmt.Errorf("Error retrieving draw function for list item %d", i),
				err,
			)
		}
		if srcRect == nil {
			return nil, srcRect, fmt.Errorf("List item %d returned nil rect", i)
		}
		if fn == nil {
			return nil, srcRect, fmt.Errorf("List item %d returned nil draw function", i)
		}
		fnList[i] = fn
		rectList[i] = srcRect
		switch list.orientation {
		case orientation_horizontal:
			expectedSize.W += srcRect.W
			if srcRect.H > expectedSize.H {
				expectedSize.H = srcRect.H
			}
		case orientation_vertical:
			expectedSize.H += srcRect.H
			if srcRect.W > expectedSize.W {
				expectedSize.W = srcRect.W
			}
		}
	}

	switch list.orientation {
	case orientation_horizontal:
		return func(dest *sdl.Rect) error {
			x := dest.X
			for i, fn := range fnList {
				itemRect := rectList[i]
				width := itemRect.W + itemRect.X
				itemDestRect := &sdl.Rect{
					X: x, Y: dest.Y,
					W: width, H: itemRect.H,
				}
				if x+width > dest.W {
					log.Warnf("list item %d is too wide %d > %d", i, x+width, dest.W)
					itemDestRect.W = 0
				}
				itemDestRect.X = x
				itemDestRect.W = width
				if err := fn(itemDestRect); err != nil {
					return err
				}
				x += width
			}
			return nil
		}, &expectedSize, nil
	case orientation_vertical:
		return func(dest *sdl.Rect) error {
			y := dest.Y
			for i, fn := range fnList {
				itemRect := rectList[i]
				height := itemRect.H + itemRect.Y
				itemDestRect := &sdl.Rect{
					X: dest.X, Y: y,
					W: itemRect.W, H: height,
				}
				if y+height > dest.H {
					log.Warnf("list item %d is too tall %d > %d", i, y+height, dest.H)
					itemDestRect.H = 0
				}
				itemDestRect.Y = y
				itemDestRect.H = height
				if err := fn(itemDestRect); err != nil {
					return err
				}
				y += height
			}
			return nil
		}, &expectedSize, nil
	default:
		return nil, &expectedSize, errors.New("Invalid orientation")
	}
}

func (text *Text) draw() (drawFn, *sdl.Rect, error) {
	var err error
	var f *ttf.Font
	if text.font == "" {
		if f, err = font.Default(); err != nil {
			return nil, defaultRect, errors.Join(
				errors.New("No font specified for text element"),
				err,
			)
		}
	} else {
		if f, err = font.Get(text.font); err != nil {
			return nil, defaultRect, errors.Join(
				fmt.Errorf("Font %s not found", text.font),
				err,
			)
		}
	}

	var surface *sdl.Surface
	// text.content is not allowed to be empty
	if text.content == "" {
		surface, err = f.RenderUTF8Solid(" ", text.color)
	} else {
		surface, err = f.RenderUTF8Solid(text.content, text.color)
	}
	if err != nil {
		return nil, defaultRect, err
	}
	defer surface.Free()

	scrRect := sdl.Rect{
		X: 0, Y: 0,
		W: surface.W, H: surface.H,
	}

	texture, err := core.Renderer().CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal(err)
	}

	return func(dest *sdl.Rect) error {
		defer texture.Destroy()
		return core.Renderer().Copy(texture, &scrRect, dest)
	}, &scrRect, nil
}
