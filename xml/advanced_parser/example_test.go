package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestID(t *testing.T) {
	var expected = map[string]interface{}{
		"title": "3",
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
