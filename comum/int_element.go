package comum

import (
	"encoding/xml"
)

type Numero int64

func (n Numero) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if n < 1 {
		return nil
	}
	e.EncodeElement(int64(n), start)
	return nil
}
