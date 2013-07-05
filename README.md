go-rss
======

Simple structs for handling RSS xml data

Install
-------

	go get github.com/vbatts/go-rss

Usage
-----

	file, _ := os.Open("ChangeLog.rss")
	dec := xml.NewDecoder(file)
	rss := rss.Rss{}
	dec.Decode(&rss)
	fmt.Printf("%#v\n", rss)

