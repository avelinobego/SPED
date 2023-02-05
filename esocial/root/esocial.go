package root

import "encoding/xml"

type ESocial struct {
	XMLName xml.Name `xml:"eSocial"`
	Data    any
}
