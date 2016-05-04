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
	url := "https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc"
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Set("User-Agent", "Holberton_School")
	req.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	//At this point the request object is: *(1) (look below)
	res, err := client.Do(req)
	//At this point the response object is: *(2)
	check(err)
	data, err := ioutil.ReadAll(res.Body)
	check(err)
	fmt.Println(string(data))
}
