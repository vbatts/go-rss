package rss

import (
	"encoding/xml"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"urn:ietf:params:xml:ns:xml atom,attr"`
	Channel Channel  `xml:"channel"`
}
