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
