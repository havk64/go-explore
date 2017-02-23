package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func fetchData(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getLocation(url string, login string, name string) map[string]string {
	var loc users
	res, err := fetchData(url)
	check(err)

	data := res.Body
	decoder := json.NewDecoder(data)
	error := decoder.Decode(&loc)
	check(error)

	defer res.Body.Close()
	obj := map[string]string{
		"location":  loc.Location,
		"full_name": name,
	}
	return obj
}

func main() {
	start := time.Now() //Starting a timer.
	ghURL := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	res, err := fetchData(ghURL)
	check(err)

	defer res.Body.Close()
	data := res.Body
	decoder := json.NewDecoder(data)
	error := decoder.Decode(&github)
	check(error)

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
	check(err)

	fmt.Println(string(ar))
	fmt.Println(time.Since(start))
}
