package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc", nil) //Creating request.
	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
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
