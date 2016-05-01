package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong!\n%v\n", e)
	}
}

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
	check(err)
	file, err := os.Create("/tmp/23")
	check(err)
	file.Chmod(0666)
	check(err)
	defer file.Close()
	w, err := file.Write(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't write the file becase of this error: %v\n", err)
	} else {
		fmt.Printf("The file was saved!\n%v bytes written", w)
	}
}
