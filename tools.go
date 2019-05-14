// tools
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func write_file(content, filepath string) {

	fw, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
	}
	defer fw.Close()
	fw.WriteString(content + "\r\n")

}

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
