package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func asyncHttpGets(user []*userObject) []*HttpResponse {
	ch := make(chan *HttpResponse, len(user)) // buffered
	responses := []*HttpResponse{}            //Empty Array of Pointers to Struct.
	for _, item := range user {
		index := item.index //Assigning variables from map obj.
		url := item.url
		login := item.login
		go func() { //Go routine
			fmt.Printf("Fetching url: %s, number: %d \n", url, index)
			data := fetchData(url)
			ch <- &HttpResponse{index, url, login, data} //Pointers to channel
		}()
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s number %d was fetched\n", r.url, r.index)
			responses = append(responses, r)
			if len(responses) == len(user) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

/* Functions to meet Sort interface */
func (a indexSorter) Len() int           { return len(a) }
func (a indexSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a indexSorter) Less(i, j int) bool { return a[i].index < a[j].index }

func main() {
	start := time.Now()
	uri := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users //Struct of Github API
	decoder := fetchData(uri)
	error := decoder.Decode(&github)
	check(error)
	defer fmt.Printf("BOOOOOMMMMM ! ! !\n30 URLs fetched in %f", time.Since(start).Seconds())
	myarray := []map[string]string{} //Initializing empty arrays
	names := []string{}              //Initializing empty arrays
	ghUser := []*userObject{}        //Initializing empty arrays of pointers.(to be used as function parameter)
	for i, item := range github.Items {
		mindex := i
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users" + "/" + login
		names = append(names, name)
		obj := &userObject{index: mindex, url: u.String(), login: login}
		ghUser = append(ghUser, obj)
	}
	results := asyncHttpGets(ghUser)
	sort.Sort(indexSorter(results))
	for _, item := range results {
		var loc users
		decoder := item.data
		error := decoder.Decode(&loc)
		check(error)
		obj := map[string]string{"location": loc.Location, "full_name": names[item.index]}
		myarray = append(myarray, obj)
	}
	ar, err := json.MarshalIndent(myarray, "", "    ")
	check(err)
	fmt.Println(string(ar))
}

/* Function fetchData to make the http requests to Github API */
func fetchData(url string) *json.Decoder {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59") //Authentication on Github API
	resp, err := client.Do(req)
	check(err)
	data := resp.Body
	decoder := json.NewDecoder(data)
	return decoder
}
