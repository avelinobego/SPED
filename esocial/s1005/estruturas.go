package s1005

import (
	"encoding/xml"
)

type EvtTabEstab struct {
	XMLName xml.Name `xml:"evtTabEstab"`
	Id      string   `xml:"id,attr"`
}
