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
func asyncHttpGets(user []*userObject) []*HttpResponse {
	ch := make(chan *HttpResponse, len(user)) // buffered
	responses := []*HttpResponse{}
	for _, item := range user {
		index := item.index
		url := item.url
		login := item.login
		go func() {
			fmt.Printf("Fetching url: %s, number: %d \n", url, index)
			data := fetchData(url)
			/*client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			check(err)
			req.Header.Set("User-Agent", "Holberton_School")
			req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
			resp, err := client.Do(req)
			data := json.NewDecoder(resp.Body)*/
			//resp.Body.Close()
			ch <- &HttpResponse{index, url, login, data}
		}() //(url)
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

func (a indexSorter) Len() int           { return len(a) }
func (a indexSorter) Swap(i, j int)      { a[i].index, a[j].index = a[j].index, a[i].index }
func (a indexSorter) Less(i, j int) bool { return a[i].index < a[j].index }

func main() {
	start := time.Now()
	uri := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	decoder := fetchData(uri)
	error := decoder.Decode(&github)
	check(error)
	defer fmt.Printf("BOOOOOMMMMM ! ! !\n30 URLs fetched in %f", time.Since(start).Seconds())
	//urls := []string{}
	myarray := []map[string]string{}
	names := []string{}
	ghUser := []*userObject{}
	for i, item := range github.Items {
		mindex := i
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users" + "/" + login
		//urls = append(urls, u.String())
		names = append(names, name)
		obj := &userObject{index: mindex, url: u.String(), login: login}
		/*obj := new(userObject)
		obj.index = mindex
		obj.login = login
		obj.url = u.String()*/
		ghUser = append(ghUser, obj)
		//obj := getLocation(u.String(), login, name)
		//myarray = append(myarray, obj)
	}
	results := asyncHttpGets(ghUser)
	//sort.Sort(indexSorter(results))
	for _, item := range results {
		var loc users
		//decoder := json.NewDecoder(item.response.Body)
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

func fetchData(url string) *json.Decoder {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	resp, err := client.Do(req)
	check(err)
	data := resp.Body
	decoder := json.NewDecoder(data)
	return decoder
}
