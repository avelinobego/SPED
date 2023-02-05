package comum

import (
	"encoding/xml"
	"time"
)

type Data time.Time

func (d Data) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(d)
	var dateString string = ""
	if t.IsZero() {
		return nil
	} else {
		dateString = time.Time(d).Format("2006-01-02")
	}
	if err := e.EncodeElement(dateString, start); err != nil {
		return err
	}
	return nil
}
