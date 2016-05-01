package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	response, err := client.Get("https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc")

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	response.Header.Set("User-Agent", "Holberton_School")
	response.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	error := ioutil.WriteFile("/tmp/23", body, 0644)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong! \n%v", error)
	}
}
