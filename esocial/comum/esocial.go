package comum

import "encoding/xml"

type ESocial[T any] struct {
	XMLName xml.Name `xml:"eSocial"`
	Data    T
}
