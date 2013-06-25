package rss

import (
	"encoding/xml"
)

type Atom struct {
	XMLName xml.Name `xml:"atom link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}
