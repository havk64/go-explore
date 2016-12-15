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

func Example() {
	selectors := map[string]interface{}{
		"title": "GoodreadsResponse book authors author id",
	}
	elements := scrape("books.xml", selectors)
	fmt.Printf("%v\n", elements)
	// Output:
	// map[title:4]

}
