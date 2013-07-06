package rss

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	// until Go1.1, to handle DefaultSpace
	//Atom           Atom     `xml:"atom link"`
	Title          string   `xml:"title"`
	Link           string   `xml:"link"`
	Description    string   `xml:"description"`
	Docs           string   `xml:"docs,omitempty"`
	Language       string   `xml:"language,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"`
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Category       []string `xml:"category,omitempty"`
	Items          []Item   `xml:"item"`
}

var ErrorNoTime = errors.New("No Time string present to parse")

/*
Attempt to parse the PubDate Field into time.Time

This expects RFC1123 or RFC1123Z.
If the pubDate field is empty or missing, then you get: time.Now, ErrorNoTime
*/
func (c *Channel) PubDateTime() (t time.Time, err error) {
	if len(c.PubDate) == 0 {
		return time.Now(), ErrorNoTime
	}
	if strings.ContainsAny(c.PubDate, "+-") {
		return time.Parse(time.RFC1123Z, c.PubDate)
	}
	return time.Parse(time.RFC1123, c.PubDate)
}
