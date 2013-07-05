package rss

type Item struct {
	Guid        string   `xml:"guid"`
	Title       string   `xml:"title"`
	Author      string   `xml:"author,omitempty"`
	PubDate     string   `xml:"pubDate"`
	Category    []string `xml:"category,omitempty"`
	Description string   `xml:"description"`
}
