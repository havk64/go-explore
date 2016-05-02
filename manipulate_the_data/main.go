package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	client := &http.Client{}

	response, err := client.Get("https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc")
	check(err)
	//if err != nil {
	//	fmt.Printf("%s", err)
	//	os.Exit(1)
	//}
	defer response.Body.Close()
	response.Header.Set("User-Agent", "Holberton_School")
	response.Header.Set("Authorization", "token 6a54def2525aa32b003337b31487e321d6a2bb59")
	//body, err := ioutil.ReadAll(response.Body)
	check(err)
	var github users //map[string]string
	decoder := json.NewDecoder(response.Body)
	error := decoder.Decode(&github)
	check(error)
	for _, item := range github.Items {
		fmt.Printf("%s\n", item.FullName)
	}

}
