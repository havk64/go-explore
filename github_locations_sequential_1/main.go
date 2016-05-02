package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type users struct {
	Items []struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		}
	}
}
type user struct {
	Location string `json:"location"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func fetchData(url string) *json.Decoder {
	client := &http.Client{}
	resp, err := client.Get(url)
	check(err)
	resp.Header.Set("User-Agent", "Holberton_School")
	resp.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	decoder := json.NewDecoder(resp.Body)
	return decoder
}

func getLocation(url string, login string, name string) map[string]string {
	var loc user
	decoder := fetchData(url)
	error := decoder.Decode(&loc)
	check(error)
	obj := map[string]string{"location": loc.Location, "full_name": name}
	return obj
}

func main() {
	uri := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	var github users
	decoder := fetchData(uri)
	error := decoder.Decode(&github)
	check(error)
	defer fmt.Println("BOOOOOMMMMM ! ! !")
	myarray := []map[string]string{}
	for _, item := range github.Items {
		name := item.FullName
		login := item.Owner.Login
		u, _ := url.Parse("https://api.github.com")
		u.Path = "/users" + "/" + login
		obj := getLocation(u.String(), login, name)
		myarray = append(myarray, obj)
	}
	ar, err := json.MarshalIndent(myarray, "", "    ")
	check(err)
	fmt.Println(string(ar))
}
