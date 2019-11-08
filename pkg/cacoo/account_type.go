package cacoo

import (
	"encoding/xml"
	"errors"
)

type AccountType string

const (
	Cacoo AccountType = "cacoo"
	Other AccountType = "other"
)

func (t *AccountType) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var e string
	err := dec.DecodeElement(&e, &start)
	if err != nil {
		return err
	}
	switch AccountType(e) {
	case Cacoo, Other:
		*t = AccountType(e)
	default:
		return errors.New("unexpected value while un-marshaling AccountType")
	}
	return nil
}
