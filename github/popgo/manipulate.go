package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type users struct {
	Items []struct {
		FullName string `json:"full_name"`
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	start := time.Now()
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc", nil) //Creating request.
	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	res, err := client.Do(req)
	check(err)
	var github users
	defer res.Body.Close()
	defer fmt.Println("Task executed in:", time.Since(start))
	defer fmt.Println("BOOOOOMMMMM ! ! !")
	decoder := json.NewDecoder(res.Body)
	error := decoder.Decode(&github)
	check(error)
	for _, item := range github.Items {
		fmt.Printf("%s\n", item.FullName)
	}
}
