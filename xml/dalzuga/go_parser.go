package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	fileBytes, err := ioutil.ReadFile("books.xml") // Read file into memory
	if err != nil {
		log.Fatal(err)
	}

	var grbq GoodReadsBookQuery

	if err = xml.Unmarshal(fileBytes, &grbq); err != nil {
		log.Fatal(err)
	}

	AuthorID := grbq.Book.Authors[0].ID

	fileBytes, err = ioutil.ReadFile("authorlistbooks.xml")
	if err != nil {
		log.Fatal(err)
	}

	var graq GoodReadsAuthorQuery

	if err := xml.Unmarshal(fileBytes, &graq); err != nil {
		log.Fatal(err)
	}

	for _, bookValue := range graq.Author.Books.Book {
		fmt.Println(bookValue.Title)
	}

	fmt.Println("start: ", graq.Author.Books.Start)
	fmt.Println("end: ", graq.Author.Books.End)
	fmt.Println("total: ", graq.Author.Books.Total)

	fmt.Println("________________________________")

	startBooks := graq.Author.Books.Start
	endBooks := graq.Author.Books.End
	totalBooks := graq.Author.Books.Total

	fmt.Println(startBooks, endBooks, totalBooks, totalBooks/endBooks)

	/* Code below is for pagination, need to code makeHTTPRequest */
	pageNumber := 1
	for totalBooks > endBooks {
		fmt.Println("_______________________REQUEST________________________")
		makeHTTPRequest("https://www.goodreads.com/author/list.xml", AuthorID, pageNumber, &graq)
		startBooks = graq.Author.Books.Start
		endBooks = graq.Author.Books.End
		totalBooks = graq.Author.Books.Total
		for _, bookValue := range graq.Author.Books.Book {
			fmt.Println(bookValue.Title)
		}
		pageNumber++
	}

	fmt.Println("Total requests:", pageNumber-1)
}

/*
 * makeHTTPRequest takes the full URL string, makes a request, and parses
 * the XML in the response into the struct pointed to by graq
 */
func makeHTTPRequest(uri string, AuthorID int, pageNumber int, graq *GoodReadsAuthorQuery) {

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

	requestBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Using more idiomatic error handling, restrict scope.
	if err = xml.Unmarshal(requestBytes, graq); err != nil {
		log.Fatal(err)
	}
}
