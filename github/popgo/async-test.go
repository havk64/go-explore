// This program fetches the 30 most popular Golang github repos(reversed sorted
// by stargazers_count) with respective locations.
// The first request gets the top repos and the second fetches the location of
// each one to be added to the final sorted result.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type data struct {
	FullName string `json:"full_name"`
	Owner    struct {
		Login string `json:"login"`
	}
	Stars int `json:"stargazers_count"`
}

type users struct {
	Location string `json:"location"`
	Items    []data
}

type mapuser struct {
	FullName string `json:"full_name"`
	Location string `json:"location"`
	Stars    int    `json:"stargazers_count"`
}

func fetchData(uri string, user *users) <-chan bool {
	ch := make(chan bool)
	go func() {
		client := &http.Client{}
		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		req.Header.Set("User-Agent", "Holberton_School")
		req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
		res, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		defer res.Body.Close()
		ch <- true
		close(ch)
	}()
	return ch
}

func getLocation(url string, item data) *mapuser {
	var loc users

	<-fetchData(url, &loc)
	return &mapuser{
		Location: loc.Location,
		FullName: item.FullName,
		Stars:    item.Stars,
	}
}

func worker(i int, item data, wg *sync.WaitGroup, result []*mapuser) {
	login := item.Owner.Login
	u, _ := url.Parse("https://api.github.com")
	u.Path = "/users/" + login
	usermap := getLocation(u.String(), item)
	result[i] = usermap
	defer wg.Done()
}

func main() {
	start := time.Now() //Starting a timer.
	fmt.Println("Starting...")

	var gh users
	var wg sync.WaitGroup
	u := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"

	<-fetchData(u, &gh)

	result := make([]*mapuser, len(gh.Items))
	for i, item := range gh.Items {
		wg.Add(1)
		go worker(i, item, &wg, result)
	}

	wg.Wait()
	ar, err := json.MarshalIndent(result, "", "    ") //Output (JSON) indented.
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	defer fmt.Println(time.Since(start))
	defer fmt.Println(string(ar))
}
