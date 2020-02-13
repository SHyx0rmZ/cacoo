package cacoo

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	s := time.Time(d).Format(time.RFC1123Z)
	return json.Marshal(s)
}

func (d Date) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(time.Time(d).Format(time.RFC1123Z), start)
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	t, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var s string
	err := dec.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	t, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
