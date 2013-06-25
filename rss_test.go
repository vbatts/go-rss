package rss

import (
	"encoding/xml"
	"fmt"
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
	fmt.Printf("%#v\n", rss)
}
