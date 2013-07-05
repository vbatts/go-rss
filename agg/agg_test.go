package agg

import (
	"log"
	"net/http"
	"testing"
  "io"
  "os"
)

func init() {
	go func() {
		http.HandleFunc("/ChangeLog.rss", func(w http.ResponseWriter, r *http.Request) {
      file, _ := os.Open("../ChangeLog.rss")
      defer file.Close()
      io.Copy(w, file)
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func TestFetch(t *testing.T) {
  r, err := FetchRss("http://localhost:8080/ChangeLog.rss")
  if err != nil {
    t.Fatalf("Failed to connect!, %s", err)
  }
  t.Logf("%#v", r)

	exp_str := "SlackBuilds.org ChangeLog"
	if r.Channel.Title != exp_str {
		t.Errorf("title [%s] did not equal %s", r.Channel.Title, exp_str)
	}
	exp_int := 29
	if len(r.Channel.Items) != exp_int {
		t.Errorf("items [%d] did not equal %d", len(r.Channel.Items), exp_int)
	}
}
