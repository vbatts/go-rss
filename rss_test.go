package rss

import (
	"encoding/xml"
	//"fmt"
	"os"
	"testing"
)

func TestThings(t *testing.T) {
	file, _ := os.Open("ChangeLog.rss")
	dec := xml.NewDecoder(file)
	rss := Rss{}
	err := dec.Decode(&rss)
	if err != nil {
		t.Errorf("%s", err)
	}
	//fmt.Printf("%#v\n", rss)
	exp_str := "SlackBuilds.org ChangeLog"
	if rss.Channel.Title != exp_str {
		t.Errorf("title [%s] did not equal %s", rss.Channel.Title, exp_str)
	}
	exp_int := 29
	if len(rss.Channel.Items) != exp_int {
		t.Errorf("items [%d] did not equal %d", len(rss.Channel.Items), exp_int)
	}
}
