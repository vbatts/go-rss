package rss

import (
	"strings"
	"time"
)

type Item struct {
	Guid        string   `xml:"guid"`
	Title       string   `xml:"title"`
	Author      string   `xml:"author,omitempty"`
	PubDate     string   `xml:"pubDate"`
	Category    []string `xml:"category,omitempty"`
	Description string   `xml:"description"`
	Link        string   `xml:"link,omitempty"`
}

// Attempt to parse the PubDate Field into time.Time
// This expects RFC1123 or RFC1123Z
func (i *Item) PubDateTime() (t time.Time, err error) {
	if strings.ContainsAny(i.PubDate, "+-") {
		return time.Parse(time.RFC1123Z, i.PubDate)
	}
	return time.Parse(time.RFC1123, i.PubDate)
}
