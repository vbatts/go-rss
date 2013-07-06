package store

import (
	"encoding/xml"
	"fmt"
	"github.com/vbatts/go-rss"
	"os"
	"testing"
)

func TestStore(t *testing.T) {
	fs, err := Open("foo.db")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("foo.db")

	// open is a second time for good measure
	fs, err = Open("foo.db")
	if err != nil {
		t.Logf("%#v\n", fs)
		t.Error(err)
	}

	urls := []string{
		"http://localhost:8080/ChangeLog.rss",
		"http://blog.hashbangbash.com/",
	}
	for _, url := range urls {
		err = fs.AddUrl(url)
		if err != nil {
			t.Errorf("Error adding URL '%s': %s", url, err)
		}
		id := fs.UrlId(url)
		if id == -1 {
			t.Errorf("rowid for %s did not get set/fetched correctly", url)
		}
	}
	if len(fs.Urls()) != 2 {
		t.Errorf("Unexpected length of Urls %#v", fs.Urls())
	}
}

func TestStoreRss(t *testing.T) {
	file, _ := os.Open("../ChangeLog.rss")
	dec := xml.NewDecoder(file)
	r := rss.Rss{}
	err := dec.Decode(&r)
	if err != nil {
		t.Errorf("%s", err)
	}
	defer file.Close()

	fs, err := Open("foo1.db")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("foo1.db")

	fs.StoreRssForUrl(r, "http://localhost:8080/ChangeLog.rss")

	fmt.Printf("%#v\n")

}
