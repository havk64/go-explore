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

type users struct {
	Location string `json:"location"`
	Items    []struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		}
	}
}

func fetchData(url string) (*json.Decoder, http.Response) {
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

	data := res.Body
	decoder := json.NewDecoder(data)
	return decoder, *res
}

func getLocation(url string, login string, name string) map[string]string {
	var loc users
	decoder, res := fetchData(url)
	err := decoder.Decode(&loc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	defer res.Body.Close()
	obj := map[string]string{
		"location":  loc.Location,
		"full_name": name,
	}
	return obj
}

func main() {
	start := time.Now() //Starting a timer.
	u := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	decoder, p := fetchData(u)
	err := decoder.Decode(&github)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	defer fmt.Println("BOOOOOMMMMM ! ! !")
	defer p.Body.Close() // Closing the http.response.Body returned as second value of fetchData().
	myarray := []map[string]string{}
	for _, item := range github.Items {
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users/" + login
		obj := getLocation(u.String(), login, name)
		fmt.Println("Fetching location of user:", login)
		myarray = append(myarray, obj)
	}

	ar, err := json.MarshalIndent(myarray, "", "    ") //Output (JSON) indented.
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	fmt.Println(string(ar))
	fmt.Println(time.Since(start))
}
