package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	client := &http.Client{}

	body := &url.Values{
		"key":  {"QaeIyNYQ484uU9WG3XxNw"},
		"id":   {"4"},
		"page": {"1"},
	}

	uri, err := url.Parse("https://www.goodreads.com/author/list.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	uri.RawQuery = body.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	xd := xml.NewDecoder(res.Body)

	for {
		tok, err := xd.Token()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Printf("%s\n", tok)
	}
	res.Body.Close()
}
