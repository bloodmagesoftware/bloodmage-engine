package ui

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/utils"
)

type Element interface {
	AppendChild(Element) error
	SetAttribute(string, string) error
	setTextContent(string) error
	setDocument(*document)
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
