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

type Image struct {
	doc *document
	id  string
	src string
}

func newImage() *Image {
	return &Image{
		doc: nil,
		id:  "",
		src: "",
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
		i.src = value
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

func (i *Image) setTextContent(content string) error {
	return errors.New("Image cannot have text content")
}
