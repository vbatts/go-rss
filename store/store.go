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

func initalize(conn *sqlite3.Conn) {
  conn.Exec("CREATE TABLE feed_urls(id INT, url TEXT)")
}

func Open(filename string) (conn *sqlite3.Conn, err error) {
  needToInitalize := true
  if FileExists(filename) {
    needToInitalize = false
  }
  conn, err = sqlite3.Open(filename)
  if err != nil {
    return conn, err
  }
  if needToInitalize {
    initalize(conn)
  }

  return
}
