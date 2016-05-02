package main

import "net/http"

type users struct {
	Location string `json:"location"`
	Items    []struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		}
	}
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}
