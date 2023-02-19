package comum

import (
	"encoding/xml"
	"time"
)

type Data time.Time

type MyStruct struct {
	Agora Data
}

func NewData(d time.Time) (result *Data) {
	result = (*Data)(&d)
	return
}

func (d Data) String() string {
	return time.Time(d).Format("2006-01")
}

func (d *Data) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d == nil {
		return nil
	}
	asTime := time.Time(*d)
	result := asTime.Format("2006-01")
	return e.EncodeElement(result, start)
}

func (d *Data) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {

	var v string
	var t time.Time
	if err = dec.DecodeElement(&v, &start); err == nil {
		if t, err = time.Parse("2006-01", v); err == nil {
			*d = Data(t)
		}
	}

	return
}
