package ui

import "errors"

type Text struct {
	doc     *document
	id      string
	content string
}

func newText() *Text {
	return &Text{
		doc:     nil,
		id:      "",
		content: "",
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
