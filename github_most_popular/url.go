package main

import (
	"net/http"
	"net/url"
)

func customURL() *url.URL {
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

	u := &url.URL{
		Scheme:   "https",
		Opaque:   "",
		User:     (*url.Userinfo)(nil),
		Host:     "api.github.com",
		Path:     "/search/repositories",
		RawPath:  "",
		RawQuery: q.Encode(),
		Fragment: "",
	}
	return u
}
func customHeader() http.Header {
	h := http.Header{
		"User-Agent": []string{
			"Holberton_School",
		},
		"Authorization": []string{
			"token 6a54def2525aa32b003337b31487e321d6a2bb59",
		},
	}
	return h
}
