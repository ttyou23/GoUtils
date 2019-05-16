// fuli
package spider

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const ROOT_URL_FULI string = "http://www.fltxt.com/"

func Fuli() {
	get_fuli_all(ROOT_URL_FULI)
}

func get_fuli_all(ori_url string) {
	get_html_content(ori_url, "body .nav a").Each(func(i int, s *goquery.Selection) {
		url := s.AttrOr("href", "")
		text := s.Text()
		fmt.Printf("title: %s - %s\n", text, url)
	})
}
