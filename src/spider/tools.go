// tools
package spider

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/tealeg/xlsx"
)

func read_excel(filepath string) {

	xlFile, err := xlsx.OpenFile(filepath)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}

func save2excel(data [][]string, filepath string) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for i, row_data := range data {
		row := sheet.AddRow()
		row.SetHeightCM(1)
		fmt.Printf("index %d  ", i)
		for j, cell_value := range row_data {
			cell := row.AddCell()
			cell.Value = cell_value
			fmt.Printf("  %d %10s  ", j, cell_value)
		}
		fmt.Printf("\n")
	}
	err = file.Save(filepath)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

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
