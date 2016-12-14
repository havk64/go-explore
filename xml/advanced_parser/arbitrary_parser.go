package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	path := "GoodreadsResponse book id"

	m := scrape("books.xml", path)
	fmt.Printf("%#v\n", m)
}

func scrape(source, p string) map[string]interface{} {
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

	path := strings.Split(p, " ")
	fmt.Printf("Path: %v\n", path)

	xd := xml.NewDecoder(file)

	res, err := parser(xd, path)
	if err != nil {
		log.Fatal(err)
	}

	result := map[string]interface{}{
		"title": res,
	}
	return result
}

func parser(xd *xml.Decoder, path []string) (string, error) {
	var result string
	last := false
	for {
		tok, err := xd.Token()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			return result, nil // break
		}

		if last {
			if test, ok := tok.(xml.CharData); ok {
				result = string(test)
				fmt.Printf("Here is what we are seeking => %s\n", test)
				return result, nil
			}
		}

		if tag, ok := tok.(xml.StartElement); ok {
			if tag.Name.Local == path[0] {
				fmt.Printf("Tag: %v\n", tag)
				if len(path) != 1 {
					path = path[1:]
				} else {
					last = true
				}
			}
		}
		// fmt.Printf("%#v\n", tok)
	}
}
