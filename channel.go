package rss

import (
	"encoding/xml"
)

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	// until Go1.1, to handle DefaultSpace
	//Atom           Atom     `xml:"atom link"`
	Title          string `xml:"title"`
	Link           string `xml:"link"`
	Docs           string `xml:"docs,omitempty"`
	Description    string `xml:"description"`
	Language       string `xml:"language,omitempty"`
	ManagingEditor string `xml:"managingEditor,omitempty"`
	WebMaster      string `xml:"webMaster,omitempty"`
	PubDate        string `xml:"pubDate,omitempty"`
	LastBuildDate  string `xml:"lastBuildDate,omitempty"`
	Generator      string `xml:"generator,omitempty"`
	Items          []Item `xml:"item"`
}
