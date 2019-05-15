// zxcs8_spider
package spider

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const ROOT_URL string = "http://www.zxcs.me/map.html"
const OUT_FILE string = "D:\\book\\zxcs8 20190514"

func Zxcs8() {

	get_zxcs8_latest(ROOT_URL)
	// get_zxcs_download("http://www.zxcs.me/post/11517")
}

func get_zxcs8_latest(ori_url string) {

	// Find the review items
	get_html_content(ROOT_URL, ".wrap #content a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		subUrl := s.AttrOr("href", "")
		// title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, subUrl)
		// get_zxcs_download(subUrl)
	})
}

func get_zxcs_download(ori_url string) { //down_2

	fmt.Println(ori_url)
	get_html_content(ori_url, "body .wrap #pleft .pagefujian .down_2 a").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i)
		// book_downurl := s.Text()
		fmt.Printf("bookurl: %s \n", s.AttrOr("href", ""))
		get_zxcs_rar(s.AttrOr("href", ""))
	})
}

func get_zxcs_rar(ori_url string) { //downfile

	fmt.Println(ori_url)
	get_html_content(ori_url, "body .wrap .content .downfile  a").Each(func(i int, s *goquery.Selection) {
		line := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("line: %s  url: %s\n", line, url)
		write_file(url, OUT_FILE+line+".txt")

	})
}
