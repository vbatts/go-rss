package agg

import (
	"errors"
	"github.com/vbatts/go-rss"
	"net/http"
  "encoding/xml"
)

var (
	ErrorResponseNotOk = errors.New("Response did not return 200 OK")
)

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
