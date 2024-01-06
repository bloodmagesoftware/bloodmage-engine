package ui

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

type Text struct {
	doc     *document
	id      string
	content string
	font    string
	color   sdl.Color
}

func newText() *Text {
	return &Text{
		doc:     nil,
		id:      "",
		content: "",
		font:    "",
		color:   sdl.Color{R: 255, G: 255, B: 255, A: 255},
	}
}

func (t *Text) AppendChild(e Element) error {
	return errors.New("Text cannot have children")
}

func (t *Text) SetAttribute(key, value string) error {
	switch key {
	case "id":
		t.id = value
	case "content":
		t.content = value
	case "font":
		t.font = value
	case "color":
        var err error
		if t.color, err = ParseColor(value); err != nil {
			return err
		}
	default:
		return errors.New("unknown attribute: " + key)
	}
	return nil
}

func (t *Text) setDocument(doc *document) {
	t.doc = doc
}

func (t *Text) Content() string {
	return t.content
}

func (t *Text) SetContent(content string) error {
	t.content = content
	return nil
}

func (t *Text) setTextContent(content string) error {
	t.content = content
	return nil
}
