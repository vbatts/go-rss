package rss

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestThingsOnSboChangelog(t *testing.T) {
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
		t.Errorf("title [%s] did not equal [%s]", rss.Channel.Title, exp_str)
	}
	exp_int := 29
	if len(rss.Channel.Items) != exp_int {
		t.Errorf("items [%d] did not equal [%d]", len(rss.Channel.Items), exp_int)
	}
	c_time, err := rss.Channel.PubDateTime()
	if err != nil {
		t.Errorf("Channel PubDateTime Parse failed with: %s", err)
	}
	fmt.Println(c_time)

	for _, item := range rss.Channel.Items {
		i_time, err := item.PubDateTime()
		if err != nil {
			t.Errorf("Item PubDateTime Parse failed with: %s", err)
		}
		fmt.Println(i_time)
	}
}

func TestThingsOnRailsPortal(t *testing.T) {
	file, _ := os.Open("portal.rss")
	dec := xml.NewDecoder(file)
	rss := Rss{}
	err := dec.Decode(&rss)
	if err != nil {
		t.Errorf("%s", err)
	}
	//fmt.Printf("%#v\n", rss)
	exp_str := "Slackware Agg Feed"
	if rss.Channel.Title != exp_str {
		t.Errorf("title [%s] did not equal [%s]", rss.Channel.Title, exp_str)
	}
	exp_int := 285
	if len(rss.Channel.Items) != exp_int {
		t.Errorf("items [%d] did not equal [%d]", len(rss.Channel.Items), exp_int)
	}

	// since this file doesn't have a Channel.PubDate, check that it works right
	now := time.Now()
	c_time, err := rss.Channel.PubDateTime()
	if err != nil && err != ErrorNoTime {
		t.Errorf("Channel PubDateTime Parse failed with: %s", err)
	}
	if err == ErrorNoTime && !c_time.After(now) {
		t.Errorf("time returned should be time.Now")
	}
	fmt.Println(c_time)

	for _, item := range rss.Channel.Items {
		i_time, err := item.PubDateTime()
		if err != nil {
			t.Errorf("Item PubDateTime Parse failed with: %s", err)
		}
		fmt.Println(i_time)
	}
}
