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

import "errors"

type Button struct {
	doc     *document
	id      string
	content Element
}

func newButton() *Button {
	return &Button{
		doc:     nil,
		id:      "",
		content: nil,
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
