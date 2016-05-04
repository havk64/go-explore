package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	client := &http.Client{}
	response, err := client.Get("https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc")
	check(err)
	defer response.Body.Close()
	response.Header.Set("User-Agent", "Holberton_School")
	response.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	body, err := ioutil.ReadAll(response.Body)
	check(err)
	fmt.Printf("%s", body)
}
