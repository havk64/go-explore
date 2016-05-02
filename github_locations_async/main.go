package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			check(err)
			req.Header.Set("User-Agent", "Holberton_School")
			req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
			resp, err := client.Do(req)
			//resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func main() {
	start := time.Now()
	uri := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	decoder := fetchData(uri)
	error := decoder.Decode(&github)
	check(error)
	defer fmt.Printf("BOOOOOMMMMM ! ! !\n30 URLs fetched in %f", time.Since(start).Seconds())
	urls := []string{}
	myarray := []map[string]string{}
	names := []string{}
	for _, item := range github.Items {
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users" + "/" + login
		urls = append(urls, u.String())
		names = append(names, name)
		//obj := getLocation(u.String(), login, name)
		//myarray = append(myarray, obj)
	}
	results := asyncHttpGets(urls)
	for i, item := range results {
		var loc users
		decoder := json.NewDecoder(item.response.Body)
		error := decoder.Decode(&loc)
		check(error)
		obj := map[string]string{"location": loc.Location, "full_name": names[i]}
		myarray = append(myarray, obj)
	}
	ar, err := json.MarshalIndent(myarray, "", "    ")
	check(err)
	fmt.Println(string(ar))
}

func fetchData(url string) *json.Decoder {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	resp, err := client.Do(req)
	check(err)
	fmt.Println(req)
	data := resp.Body
	//defer resp.Body.Close()
	decoder := json.NewDecoder(data)
	return decoder
}
