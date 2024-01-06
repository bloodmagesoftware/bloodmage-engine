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
