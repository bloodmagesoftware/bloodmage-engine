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
