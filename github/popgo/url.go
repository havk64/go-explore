package main

import (
	"net/http"
	"net/url"
)

func customURL() *url.URL {
	q := &url.Values{ // Custom Query created based in the the type url.Values,
		"q":     {"language:go"},
		"sort":  {"stars"},
		"order": {"desc"},
	}
	/* We return the reference to url.URL struct, which has the right format to *
	 * stringify the Url using the encoded reference to url.Values as RawQuery. */
	return &url.URL{
		Scheme:   "https",
		Host:     "api.github.com",
		Path:     "/search/repositories",
		RawQuery: q.Encode(), //Using the custom Query created above.
	}
}

/* Using referent to http.Header Struct to customize our request Header.      */
func customHeader() http.Header {
	return http.Header{
		"User-Agent": {"Holberton_School"},
		"Authorization": {
			"token 9f5400d5762cdca712e84b8c921cfa801c3dbeb6"},
	}
}
