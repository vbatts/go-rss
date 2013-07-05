package store

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"os"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

type FeedStore struct {
	conn *sqlite3.Conn
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

var (
	Schema = []string{
		"CREATE TABLE feed_urls(url TEXT)",
		"CREATE TABLE feed_info(feed_id INT, title TEXT, lastBuildDate INT, description TEXT, link TEXT)",
		"CREATE TABLE feed_items(feed_id INT, guid TEXT, pubDate TEXT, title TEXT, description TEXT)",
	}
)

func initalize(conn *sqlite3.Conn) error {
	for _, query := range Schema {
		err := conn.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

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
