// This program fetch the top Golang Github repositories and writes the results
// to a file. Using basic concurrency through Goroutines and Channels.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	body := request()

	file, err := os.Create("/tmp/23")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	w, err := file.Write(<-body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	} else {
		fmt.Printf("The file was saved!\n%v bytes written\n", w)
	}

	defer fmt.Printf("%v\n", time.Since(start))
	defer file.Close()
}

func request() <-chan []byte {
	ch := make(chan []byte)
	go func() {
		url := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
		client := &http.Client{}

		req, err := http.NewRequest("GET", url, nil) //Creating request.
		if err != nil {
			fmt.Fprintf(os.Stderr, "write2disk: %v\n", err)
		}

		req.Header.Set("User-Agent", "Holberton_School")
		req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")

		res, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "write2disk: %v\n", err)
			os.Exit(1)
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "write2disk: %v\n", err)
		}
		ch <- body
		close(ch)
	}()
	return ch
}
