package fetch

import (
	"encoding/xml"
	"errors"
	"github.com/vbatts/go-rss"
	"net/http"
	"time"
)

var (
	ErrorResponseNotOk = errors.New("Response did not return 200 OK")
)

func RemoteIsNewer(urlStr string, t time.Time) (isNewer bool, err error) {
	resp, err := http.Head(urlStr)
	if err != nil {
		return true, err
	}

	// if the remote hasn't set Last-Modified header, assume that it is newer
	// and that it'll need to be refetched.
	if len(resp.Header.Get("Last-Modified")) == 0 {
		return true, nil
	}

	remote_t, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		return true, err
	}
	if remote_t.After(t) {
		return true, nil
	}

	return false, nil
}

func FetchRss(urlStr string) (r *rss.Rss, err error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, ErrorResponseNotOk
	}

	dec := xml.NewDecoder(resp.Body)
	r = &rss.Rss{}
	err = dec.Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
