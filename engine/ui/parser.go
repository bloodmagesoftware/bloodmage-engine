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
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Element interface {
	AppendChild(Element) error
	SetAttribute(string, string) error
	setTextContent(string) error
	setDocument(*document)
	draw() (drawFn, *sdl.Rect, error)
}

func CreateElement(name string) (Element, error) {
	switch name {
	case "List":
		return newList(), nil
	case "Button":
		return newButton(), nil
	case "Text":
		return newText(), nil
	case "Image":
		return newImage(), nil
	default:
		return nil, fmt.Errorf("unknown element type: %s", name)
	}
}

func Parse(path string) (*document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	var root Element
	doc := newDocument()
	elStack := utils.NewStack[Element]()

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch t := token.(type) {
		case xml.StartElement:
			// create element
			el, err := CreateElement(t.Name.Local)
			if err != nil {
				return nil, err
			}
			// might be root element
			if root == nil {
				root = el
				doc.root = el
			}
			el.setDocument(&doc)
			// set attributes
			for _, attr := range t.Attr {
				if attr.Name.Space != "" {
					continue
				}
				if attr.Name.Local == "id" {
					doc.idMap[attr.Value] = el
				}
				if err := el.SetAttribute(attr.Name.Local, attr.Value); err != nil {
					return nil, errors.Join(
						fmt.Errorf("error setting attribute %s=%s on element %s", attr.Name.Local, attr.Value, t.Name.Local),
						err,
					)
				}
			}
			// append to parent if exists
			if parent, hasParent := elStack.Peek(); hasParent {
				if err := (*parent).AppendChild(el); err != nil {
					return nil, err
				}
			}
			// push to stack for future children
			elStack.Push(el)

		case xml.CharData:
			text := strings.TrimSpace(string(t))
			if text == "" {
				continue
			}
			if el, hasElement := elStack.Peek(); hasElement {
				if err := (*el).setTextContent(text); err != nil {
					return nil, err
				}
			}

		case xml.EndElement:
			// pop from stack because we're done with this element
			elStack.Pop()
		}
	}

	if root == nil {
		return nil, fmt.Errorf("no root element found")
	}

	return &doc, nil
}
