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
}

type users struct {
	Location string `json:"location"`
	Items    []data
}

func fetchData(url string, user *users) <-chan bool {
	ch := make(chan bool, 1)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
	return ch
}

func getLocation(url string, name string) map[string]string {
	var loc users
	<-fetchData(url, &loc)

	usermap := map[string]string{
		"location":  loc.Location,
		"full_name": name,
	}
	return usermap
}

func main() {
	fmt.Println("Starting...")
	start := time.Now() //Starting a timer.
	u := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	fetchData(u, &github)

	result := make([]map[string]string, len(github.Items))
	for i, item := range github.Items {
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users/" + login
		usermap := getLocation(u.String(), name)
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
