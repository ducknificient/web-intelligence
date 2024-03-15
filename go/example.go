package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func test_crawling() {

	var seed_url = `https://scrapeme.live/shop/`

	seed_url = `http://metalsucks.net`

	// download the target HTML document
	res, err := http.Get(seed_url)
	if err != nil {
		log.Fatal("Failed to connect to the target page", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	// convert the response buffer to bytes
	/*
		byteBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal("Error while reading the response buffer", err)
		}
	*/

	// convert the byte HTML content to string and
	// print it
	// html := string(byteBody)
	// fmt.Println(html)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

}
