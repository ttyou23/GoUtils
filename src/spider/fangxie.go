// fangxie
package spider

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const ROOT_URL_FANGXIE string = "https://www.cdfangxie.com/"

func Fangxie() {
	fangxie_sort()
}

func fangxie_sort() {
	get_html_content(ROOT_URL_FANGXIE, "body div[class=cont1_rukou] a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("title: %s - %s\n", text, url)
		fangxie_info(ROOT_URL_FANGXIE + url)
		return false
	})
}

func fangxie_info(ori_url string) {
	selection := get_html_content(ori_url, "body .wrapper .main .right_cont")
	selection.Find(".ul_list li a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("info: %s - %s\n", title, url)
		return true
	})

	selection.Find(".pages2 a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("sort: %s - %s\n", title, url)
		return true
	})
}
