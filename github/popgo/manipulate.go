package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type ghUser struct {
	Items []struct {
		FullName string `json:"full_name"`
	}
}

func main() {
	start := time.Now()
	var user ghUser
	<-request(&user)

	defer fmt.Println("Task executed in:", time.Since(start))
	for _, item := range user.Items {
		fmt.Printf("%s\n", item.FullName)
	}
}

func request(user *ghUser) <-chan bool {
	ch := make(chan bool)
	go func() {
		uri := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
		client := &http.Client{}

		req, err := http.NewRequest("GET", uri, nil) //Creating request.
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}

		req.Header.Set("User-Agent", "Holberton_School")
		req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")

		res, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}

		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		ch <- true
		close(ch)
	}()
	return ch
}
