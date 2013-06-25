package rss

type Item struct {
	Guid        string `xml:"guid"`
	Title       string `xml:"title"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}
