package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	xmlFile := flag.String("file", "books.xml", "XML source file to be decoded")
	path := flag.String("path", "GoodreadsResponse book authors author id", "Path to search in xml file")
	flag.Parse()

	defSelector := map[string]interface{}{"title": *path}
	m := scrape(*xmlFile, defSelector)
	fmt.Printf("%#v\n", m)
}

func scrape(source string, selector map[string]interface{}) map[string]interface{} {
	p := selector["title"].(string)
	path := strings.Split(p, " ")
	fmt.Printf("Path: %v\n", path)

	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

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
			if content, ok := tok.(xml.CharData); ok {
				result = string(content)
				// fmt.Printf("Here is what we want => %s\n", test)
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
