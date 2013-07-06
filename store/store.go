package store

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/vbatts/go-rss"
	"os"
)

var (
	Schema = []string{
		"CREATE TABLE feed_urls(url TEXT)",
		"CREATE TABLE feed_info(feed_id INT, title TEXT, lastBuildDate TEXT, description TEXT, link TEXT)",
		"CREATE TABLE feed_items(feed_id INT, guid TEXT, pubDate TEXT, title TEXT, description TEXT)",
	}
)

// Not really a data store function, but a simple os check for whether
// filename already exists
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
Container for the given open data store

This is used for storing and fetching whole rss data structures.
Easy to store once retrieved, and easy to query for use.
*/
type FeedStore struct {
	conn *sqlite3.Conn
}

// Store an rss.Rss structure away, given it's the rss URL's urlId
func (fs *FeedStore) StoreRssForUrlId(r rss.Rss, urlId int64) error {
	args := sqlite3.NamedArgs{
		"$feed_id":       urlId,
		"$title":         r.Channel.Title,
		"$lastBuildDate": r.Channel.LastBuildDate,
		"$description":   r.Channel.Description,
		"$link":          r.Channel.Link,
	}
	if fs.infoExistsForUriId(urlId) {
		fs.conn.Exec("DELETE FROM feed_info WHERE feed_id = ?", urlId)
	}
	err := fs.conn.Exec("INSERT INTO feed_info VALUES($feed_id, $title, $lastBuildDate, $description, $link)", args)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FeedStore) infoExistsForUriId(urlId int64) bool {
	var count int64
	query := "SELECT count(1) FROM feed_info WHERE feed_id = ?"
	for s, err := fs.conn.Query(query, urlId); err == nil; err = s.Next() {
		s.Scan(&count)
	}
	return count > 0
}

// Store an rss.Rss structure away, given it's the rss URL's string
func (fs *FeedStore) StoreRssForUrl(r rss.Rss, urlStr string) error {
	id := fs.UrlId(urlStr)
	if id == -1 {
		fs.AddUrl(urlStr)
		id = fs.UrlId(urlStr)
	}
	return fs.StoreRssForUrlId(r, id)
}

// Insert urlStr URL into the FeedStore
func (fs *FeedStore) AddUrl(urlStr string) error {
	return fs.conn.Exec("INSERT INTO feed_urls VALUES(?)", urlStr)
}

// Get the ID for a given URL
func (fs *FeedStore) UrlId(urlStr string) int64 {
	query := "SELECT rowid FROM feed_urls WHERE url = ?"
	for s, err := fs.conn.Query(query, urlStr); err == nil; err = s.Next() {
		var rowid int64
		s.Scan(&rowid)
		return rowid
	}
	return -1
}

// Get a map of stored URLs
func (fs *FeedStore) Urls() (urls map[int64]string) {
	urls = make(map[int64]string)
	query := "SELECT rowid, * FROM feed_urls"
	//row := make(sqlite3.RowMap)
	for s, err := fs.conn.Query(query); err == nil; err = s.Next() {
		var (
			rowid int64
			url   string
		)
		//s.Scan(&id, row)
		s.Scan(&rowid, &url)
		urls[rowid] = url
	}

	return urls
}

func initalize(conn *sqlite3.Conn) error {
	for _, query := range Schema {
		err := conn.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

// Creates and initalizes, or Reopens a database store for RSS Feeds
func Open(filename string) (fs *FeedStore, err error) {
	needToInitalize := true
	if FileExists(filename) {
		needToInitalize = false
	}
	conn, err := sqlite3.Open(filename)
	if err != nil {
		return nil, err
	}
	if needToInitalize {
		err = initalize(conn)
		if err != nil {
			return nil, err
		}
	}
	fs = &FeedStore{conn}

	return fs, nil
}
