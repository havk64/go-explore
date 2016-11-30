package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	body := request()

	file, err := os.Create("/tmp/23")
	if err != nil {
		log.Fatal(err)
	}

	defer fmt.Printf("%v\n", time.Since(start))
	defer file.Close()
	w, err := file.Write(body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	} else {
		fmt.Printf("The file was saved!\n%v bytes written\n", w)
	}
}

func request() []byte {
	url := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil) //Creating request.
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
