package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	// Get the Author ID(in this case it's going to return 4)
	AuthorID, err := getAuthorID("books.xml")
	if err != nil {
		log.Fatal(err)
	}
	// Create the structure that it's going to get the result
	results := make(map[int][]string)
	// Paginate the books up to the point where we have all the books
	// The condition is empty(between the two ";")
	for pageNumber := 1; ; pageNumber++ {
		var graq GoodReadsAuthorQuery
		// Worker returns a bool that indicates if the last book is equal total books
		// if "stops" is true it breaks and exits the loop
		// Notice the syntax: We have an assignment inside the if statement
		// Stops gets a bool value and if it's true it "break(s)"
		if stops := worker(AuthorID, pageNumber, &graq, results); stops {
			break
		}
	}
	// Iterate through the results printing it
	for i, v := range results {
		fmt.Printf("_______________________PAGE %v________________________\n", i+1)
		for i, s := range v {
			fmt.Println(i+1, s)
		}
	}
}

func getAuthorID(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	var v GoodReadsBookQuery

	if err := xml.NewDecoder(file).Decode(&v); err != nil {
		log.Fatal(err)
	}

	AuthorID := v.Book.Authors[0].ID
	return AuthorID, nil
}

// makeHTTPRequest takes the full URL string, makes a request, and parses
// the XML in the response into the struct pointed to by graq
func makeHTTPRequest(AuthorID, pageNumber int, graq *GoodReadsAuthorQuery) {
	uri := "https://www.goodreads.com/author/list.xml"
	client := &http.Client{}

	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("key", `kDkKnUxiz8cRBJhVjrtSA`)
	q.Set("id", strconv.Itoa(AuthorID))
	q.Set("page", strconv.Itoa(pageNumber))
	u.RawQuery = q.Encode()
	fullURL := u.String()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if err = xml.NewDecoder(resp.Body).Decode(graq); err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}

func worker(AuthorID, pageNumber int, graq *GoodReadsAuthorQuery, results map[int][]string) bool {
	makeHTTPRequest(AuthorID, pageNumber, graq)

	r := make([]string, len(graq.Author.Books.Book))

	for i, bookValue := range graq.Author.Books.Book {
		r[i] = bookValue.Title
	}

	results[pageNumber-1] = r

	startBooks := graq.Author.Books.Start
	endBooks := graq.Author.Books.End
	totalBooks := graq.Author.Books.Total
	fmt.Printf("Page %v => ", pageNumber)
	fmt.Printf("Start: %v, End: %v, Total: %v\n", startBooks, endBooks, totalBooks)

	if endBooks == totalBooks {
		return true // stops the loop
	}
	return false // keep iterating through pages
}
