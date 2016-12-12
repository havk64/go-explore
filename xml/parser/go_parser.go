package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	AuthorID, err := getAuthorID("books.xml")
	if err != nil {
		log.Fatal(err)
	}

	graq, err := parseAuthorBooks("authorlistbooks.xml")
	if err != nil {
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

	var wg sync.WaitGroup
	results := make(map[int][]string)
	/* Code below is for pagination, need to code makeHTTPRequest */
	pageNumber := 1
	for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", pageNumber)
		wg.Add(1)
		go func(wg *sync.WaitGroup, AuthorID, pageNumber int, graq *GoodReadsAuthorQuery) {
			makeHTTPRequest(AuthorID, pageNumber, graq)

			r := make([]string, len(graq.Author.Books.Book))

			for i, bookValue := range graq.Author.Books.Book {
				r[i] = bookValue.Title
			}
			results[pageNumber-1] = r
			defer wg.Done()
		}(&wg, AuthorID, pageNumber, graq)
		pageNumber++
	}

	wg.Wait()
	for _, v := range results {
		fmt.Println("_______________________REQUEST________________________")
		for _, s := range v {
			fmt.Println(s)
		}
	}
	fmt.Println("Total requests:", pageNumber-1)
}

/*
 * makeHTTPRequest takes the full URL string, makes a request, and parses
 * the XML in the response into the struct pointed to by graq
 */
func makeHTTPRequest(AuthorID int, pageNumber int, graq *GoodReadsAuthorQuery) {
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

	// Using more idiomatic error handling, restrict scope.
	// Using NewDecoder allows us to use resp.Body directly without ioutil
	// Because res.Body is an *http.Response which satisfy io.Reader interface
	if err = xml.NewDecoder(resp.Body).Decode(graq); err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
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

func parseAuthorBooks(f string) (*GoodReadsAuthorQuery, error) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	var v GoodReadsAuthorQuery

	if err := xml.NewDecoder(file).Decode(&v); err != nil {
		log.Fatal(err)
	}

	return &v, nil
}
