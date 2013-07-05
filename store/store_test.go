package store

import (
	"fmt"
	"testing"
  "os"
)

func TestStore(t *testing.T) {
	c, err := Open("foo.db")
	if err != nil {
		t.Error(err)
	}
  defer os.Remove("foo.db")
	fmt.Printf("%#v\n", c)

  // open is a second time for good measure
	c, err = Open("foo.db")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", c)
}
