package main

import (
	"fmt"
	"net/url"
)

func main() {
	q := &url.Values{
		"q": []string{
			"language:go",
		},
		"sort": []string{
			"stars",
		},
		"order": []string{
			"desc",
		},
	}

	t := &url.URL{
		Scheme:   "https",
		Opaque:   "",
		User:     (*url.Userinfo)(nil),
		Host:     "api.github.com",
		Path:     "/search/repositories",
		RawPath:  "",
		RawQuery: q.Encode(),
		Fragment: "",
	}

	fmt.Println(t)
	fmt.Println(t.Query())
	fmt.Println(q.Encode())

}
