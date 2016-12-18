package xmlParser

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"
)


func scrape(source string, selector map[string]interface{}) map[string]interface{} {
	p, ok := selector["title"].(string)
	if !ok {
		log.Fatal("The value of \"title\" is expected to be of type \"string\"\n")
	}

	path := strings.Split(p, " ")

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
			return result, nil
		}

		if last {
			if content, ok := tok.(xml.CharData); ok {
				result = string(content)
				return result, nil
			}
		}

		if tag, ok := tok.(xml.StartElement); ok {
			if tag.Name.Local == path[0] {
				if len(path) != 1 {
					path = path[1:]
				} else {
					last = true
				}
			}
		}
	}
}
