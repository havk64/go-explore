package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestID(t *testing.T) {
	var expected = map[string]interface{}{
		"title": "4",
	}

	selectors := map[string]interface{}{
		"title": "GoodreadsResponse book authors author id",
	}

	elements := scrape("books.xml", selectors)
	eq := reflect.DeepEqual(elements, expected)
	if !eq {
		t.Errorf("Expected: %v, got: %v", expected, elements) //("Scrape(%s) = %v, want %v", elements, expected)
	}
}

func TestBooks(t *testing.T) {
	var expected = map[string]interface{}{
		"title": "The Hitchhiker's Guide to the Galaxy (Hitchhiker's Guide to the Galaxy, #1)",
	}

	selectors := map[string]interface{}{
		"title": "GoodreadsResponse author books book title",
	}

	elements := scrape("authorlistbooks.xml", selectors)
	eq := reflect.DeepEqual(elements, expected)
	if !eq {
		t.Errorf("Expected: %v, got: %v", expected, elements) //("Scrape(%s) = %v, want %v", elements, expected)
	}
}

func Example() {
	selectors := map[string]interface{}{
		"title": "GoodreadsResponse book authors author id",
	}
	elements := scrape("books.xml", selectors)
	fmt.Printf("%v\n", elements)
	// Output:
	// map[title:4]
}

func TestParser(t *testing.T) {
	expected := "4"

	file, err := os.Open("books.xml")
	if err != nil {
		log.Fatal(err)
	}

	xd := xml.NewDecoder(file)
	path := []string{"GoodreadsResponse", "book", "authors", "author"}

	str, err := parser(xd, path)
	if err != nil {
		log.Fatal(err)
	}

	if str != expected {
		t.Errorf("Expected: %#v, Got: %#v\n", expected, str)
	}
}

/*
func TestEvolved(t *testing.T) {
	selectors := map[string]interface{}{
		"title":   "Movies Movie Title",
		"release": "Movies Movie Release",
		"actors": map[string]string{
			"actor_name":      "Movies Movie Cast Actor Name",
			"actor_character": "Movies Movie Cast Actor Character",
		},
	}
	elements := scrape("books.xml", selectors)
	fmt.Println(elements)
}
*/
