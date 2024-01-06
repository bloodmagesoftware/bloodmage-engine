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
