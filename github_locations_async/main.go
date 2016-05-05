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
func asyncHTTPGets(user []*userObject) []*HTTPResponse {
	ch := make(chan *HTTPResponse, len(user)) // buffered
	responses := []*HTTPResponse{}            //Empty Array of Pointers to Struct.
	for _, item := range user {
		index := item.index //Assigning variables from map obj.
		url := item.url
		login := item.login
		go func() { //Goroutine
			fmt.Printf("Fetching url: %s, ranking: %d \n", url, index+1)
			data, res := fetchData(url)
			ch <- &HTTPResponse{index, url, login, data, res} //Pointers to channel
		}()
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s, ranking %d was fetched\n", r.url, r.index+1)
			responses = append(responses, r)
			if len(responses) == len(user) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

/* Functions to meet Sort interface requirements */
func (a indexSorter) Len() int           { return len(a) }
func (a indexSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a indexSorter) Less(i, j int) bool { return a[i].index < a[j].index }

func main() {
	start := time.Now()
	u := customURL() // Replacing previous code below with this function.
	/* ====================================================================================
	  u, error := url.Parse("https://api.github.com/search/repositories") //Parsing the URL
		check(error)
		q := u.Query() //Getting url.Query() in order to specify the Query
		/* This could work too(instead of adding each item to the url.Query() objc. as in next lines):
		q = map[string][]string{"q": []string{"language:go"}, "sort": []string{"stars"}, "order": []string{"desc"}}
		>> Next step is to use a struct to build the Query.
		q.Add("q", "language:go") // Search for all Golang repositories
		q.Add("sort", "stars")    // Sort by amount of Stars
		q.Add("order", "desc")    // In descentant order
		u.RawQuery = q.Encode()   // Encoding the query to make it a encode string
	* ===================================================================================== */
	github := &users{} // Struct of Github API // Before was declared as: `var github users`.
	decoder, res := fetchData(u.String())
	err := decoder.Decode(&github)
	check(err)
	defer res.Body.Close()
	defer fmt.Println("BOOOOOMMMMM ! ! !\n30 URLs fetched in ", time.Since(start))
	myarray := []map[string]interface{}{} //Initializing empty arrays
	names := []string{}                   //Initializing empty arrays
	ghUser := []*userObject{}             //Initializing empty arrays of pointers.(to be used as function parameter)
	for i, item := range github.Items {
		index := i
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users/" + login
		names = append(names, name)
		obj := &userObject{index: index, url: u.String(), login: login}
		ghUser = append(ghUser, obj)
	}
	results := asyncHTTPGets(ghUser)
	sort.Sort(indexSorter(results)) //==> Using sort in the Index in order to output in the same order first request.
	for _, item := range results {
		loc := &users{} // Before was declared: `var loc users`.
		decoder := item.data
		error := decoder.Decode(&loc)
		check(error)
		defer item.res.Body.Close()
		/* Object to be displayed in the output */
		obj := map[string]interface{}{"location": loc.Location, "full_name": names[item.index], "ranking": (item.index + 1)}
		myarray = append(myarray, obj)
	}
	ar, err := json.MarshalIndent(myarray, "", "    ") /* Indenting the output(Json Prettifyied) */
	check(err)
	fmt.Println(string(ar))
}

/* Function fetchData to make the http requests to Github API */
func fetchData(url string) (*json.Decoder, http.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header = customHeader()
	resp, err := client.Do(req)
	check(err)
	data := resp.Body
	decoder := json.NewDecoder(data) // Parsing the JSON Object.
	return decoder, *resp
}
