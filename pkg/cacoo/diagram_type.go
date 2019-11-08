package cacoo

import (
	"encoding/xml"
	"errors"
)

type DiagramType string

const (
	Normal   DiagramType = "normal"
	Stencil  DiagramType = "stencil"
	Template DiagramType = "template"
)

func (t *DiagramType) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var e string
	err := dec.DecodeElement(&e, &start)
	if err != nil {
		return err
	}
	switch DiagramType(e) {
	case Normal, Stencil, Template:
		*t = DiagramType(e)
	default:
		return errors.New("unexpected value while un-marshaling DiagramType")
	}
	return nil
}

func (t DiagramType) String() string {
	return string(t)
}
