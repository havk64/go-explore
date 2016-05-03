package main

import "encoding/json"

type users struct {
	Location string `json:"location"`
	Items    []struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		}
	}
}

//HTTPResponse struct is used to get the make the async requests." */
type HTTPResponse struct {
	index int
	url   string
	login string
	data  *json.Decoder
	//err   error
}
type userObject struct {
	index int
	url   string
	login string
}

type indexSorter []*HTTPResponse
