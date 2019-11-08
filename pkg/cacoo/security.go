package cacoo

import (
	"encoding/xml"
	"errors"
)

type SecurityType string

const (
	Private SecurityType = "private"
	URL     SecurityType = "url"
	Public  SecurityType = "public"
)

func (t *SecurityType) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var e string
	err := dec.DecodeElement(&e, &start)
	if err != nil {
		return err
	}
	switch SecurityType(e) {
	case Private, URL, Public:
		*t = SecurityType(e)
	default:
		return errors.New("unexpected value while un-marshaling SecurityType")
	}
	return nil
}
