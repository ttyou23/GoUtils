// tools
package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func get_html_content(ori_url, selector string) *goquery.Selection {
	// Request the HTML page.
	res, err := http.Get(ori_url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	return doc.Find(selector)
}
