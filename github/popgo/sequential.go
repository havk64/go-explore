// This program fetches most popular Go repositories on github
// and parses its json response sequentially, not concurrently.
// Go-explore is a sequence of experiments with the Go Language.
// Here I'm exploring some godoc features to format documentation.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
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
	Items    []data // a collection of struct "data"
}

type usermap struct {
	FullName string `json:"full_name"`
	Location string `json:"location"`
	Stars    int    `json:"stars"`
}

// fetchData accepts a url and a pointer to "users" struct, fetches the url,
// decode the json response using "users" struct and return a bool
// for syncronization.
func fetchData(url string, user *users) <-chan bool {
	ch := make(chan bool, 1)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 9f5400d5762cdca712e84b8c921cfa801c3dbeb6")

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
	return ch
}

// getLocation accepts a url and github username, uses fetchData to fetch the
// location of each user and returns a map with the result.
func getLocation(url string, item data) *usermap {
	var loc users
	name := item.FullName
	stars := item.Stars
	<-fetchData(url, &loc)

	usermap := &usermap{
		FullName: name,
		Location: loc.Location,
		Stars:    stars,
	}
	return usermap
}

func main() {
	fmt.Println("Starting...")
	start := time.Now() //Starting a timer.
	u := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	fetchData(u, &github)

	result := make([]*usermap, len(github.Items))
	for i, item := range github.Items {
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users/" + login
		usermap := getLocation(u.String(), item)
		fmt.Println("Fetching location of user:", login)
		result[i] = usermap
	}

	ar, err := json.MarshalIndent(result, "", "    ") //Output (JSON) indented.
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	fmt.Println(string(ar))
	fmt.Println(time.Since(start))
}
