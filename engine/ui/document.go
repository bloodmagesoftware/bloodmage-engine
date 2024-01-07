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

type document struct {
	root  Element
	idMap map[string]Element
}

func newDocument() document {
	return document{
		idMap: make(map[string]Element),
	}
}

func (d *document) RootElement() Element {
	return d.root
}

func (d *document) GetElementById(id string) (Element, bool) {
	e, ok := d.idMap[id]
	return e, ok
}

func (d *document) GetTextElementById(id string) (*Text, bool) {
	e, ok := d.idMap[id]
	if !ok {
		return nil, false
	}
	t, ok := e.(*Text)
	if !ok {
		return nil, false
	}
	return t, true
}

func (d *document) GetButtonElementById(id string) (*Button, bool) {
	e, ok := d.idMap[id]
	if !ok {
		return nil, false
	}
	b, ok := e.(*Button)
	if !ok {
		return nil, false
	}
	return b, true
}

func (d *document) GetImageElementById(id string) (*Image, bool) {
	e, ok := d.idMap[id]
	if !ok {
		return nil, false
	}
	i, ok := e.(*Image)
	if !ok {
		return nil, false
	}
	return i, true
}

func (d *document) GetListElementById(id string) (*List, bool) {
	e, ok := d.idMap[id]
	if !ok {
		return nil, false
	}
	l, ok := e.(*List)
	if !ok {
		return nil, false
	}
	return l, true
}
