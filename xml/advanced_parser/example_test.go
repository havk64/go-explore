package main

import "fmt"

func Example() {
	selectors := map[string]interface{}{
		"title": "GoodreadsResponse book authors author id",
	}
	elements := scrape("books.xml", selectors)
	fmt.Printf("%v\n", elements)
	// Output:
	// map[title:4]
}
