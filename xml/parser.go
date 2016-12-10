package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("books.xml")
	if err != nil {
		log.Fatal(err)
	}

	xd := xml.NewDecoder(file)

	for {
		tok, err := xd.Token()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Printf("%q\n", tok)
	}
}
