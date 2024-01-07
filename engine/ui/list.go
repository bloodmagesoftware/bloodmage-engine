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

import "errors"

type orientation uint8

const (
	orientation_unset orientation = iota
	orientation_vertical
	orientation_horizontal
)

type List struct {
	doc *document
	orientation
	items []Element
}

func newList() *List {
	return &List{
		doc:         nil,
		orientation: orientation_unset,
		items:       nil,
	}
}

func (l *List) AppendChild(e Element) error {
	l.items = append(l.items, e)
	return nil
}

func (l *List) SetAttribute(key, value string) error {
	switch key {
	case "orientation":
		switch value {
		case "vertical":
			l.orientation = orientation_vertical
		case "horizontal":
			l.orientation = orientation_horizontal
		default:
			return errors.New("unknown orientation: " + value)
		}
	default:
		return errors.New("unknown attribute: " + key)
	}
	return nil
}

func (l *List) setDocument(doc *document) {
	l.doc = doc
}

func (l *List) Items() []Element {
	return l.items
}

func (l *List) setTextContent(content string) error {
	return errors.New("list cannot have text content")
}
