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

func main() {
	start := time.Now()
	response := <-request()

	defer fmt.Println("Task executed in:", time.Since(start))
	for _, item := range response.Items {
		fmt.Printf("%s\n", item.FullName)
	}
}

func request() <-chan users {
	ch := make(chan users)
	var gh users
	go func() {
		uri := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
		client := &http.Client{}

		req, err := http.NewRequest("GET", uri, nil) //Creating request.
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "Holberton_School")
		req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		// body, err := ioutil.ReadAll(res.Body)
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&gh)
		if err != nil {
			log.Fatal(err)
		}
		ch <- gh
		close(ch)
	}()
	return ch
}
