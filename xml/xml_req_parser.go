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
	// Create the http client
	client := &http.Client{}

	// Build the query string
	body := &url.Values{
		"key":  {"QaeIyNYQ484uU9WG3XxNw"},
		"id":   {"4"},
		"page": {"1"},
	}

	// Parse the Url
	uri, err := url.Parse("https://www.goodreads.com/author/list.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	// Put the encoded string in the url
	uri.RawQuery = body.Encode()

	// Build the http request
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	// Send the http request and get the result
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v", err)
	}

	// Parse the result
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
	//Close the response body
	res.Body.Close()
}
