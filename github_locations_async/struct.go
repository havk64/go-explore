package main

import (
	"encoding/json"
	"net/http"
)

type users struct {
	Location string `json:"location"`
	Items    []struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		}
	}
}

// The HTTPResponse struct is used to get the make the async requests."
type HTTPResponse struct {
	index int
	url   string
	login string
	data  *json.Decoder
	res   http.Response
	//err   error
}
type userObject struct {
	index int
	url   string
	login string
}

type indexSorter []*HTTPResponse
