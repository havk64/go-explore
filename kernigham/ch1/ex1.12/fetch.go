// http client to fetch and test the Server-lissajous exercise created previously
// Parameters can be added to url.Values
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	client := &http.Client{}
	uri, err := url.Parse("http://localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing: %v\n", err)
	}

	body := &url.Values{
		"cycles": {"20"},
	}

	uri.RawQuery = body.Encode()
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		os.Exit(1)
	}

	n, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying response body: %v\n", err)
	}

	fmt.Printf("%#v\n", n)
	defer res.Body.Close()
}
