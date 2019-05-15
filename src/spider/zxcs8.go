// zxcs8
package spider

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const ROOT_URL_ZXCS8 string = "http://www.zxcs.me/map.html"
const OUT_FILE string = "D:\\book\\zxcs8 20190515"

func Zxcs8() {

	zxcs8_get_latest()
	// get_zxcs_all()

	//测试-----
	// get_zxcs_download("http://www.zxcs.me/post/11517")
	// get_zxcs_sort("http://www.zxcs.me/sort/55")
}

func zxcs_get_all() {

	//下载所有书籍
	get_html_content(ROOT_URL_ZXCS8, "body .wrap #sort ul a:contains('(')").EachWithBreak(func(i int, s *goquery.Selection) bool {
		band := s.Text()
		subUrl := s.AttrOr("href", "")
		// title := s.Find("i").Text()
		fmt.Printf("record %d: %s - %s\n", i, band, subUrl)
		zxcs_get_sort(subUrl)
		return false
	})
}

func zxcs_get_sort(ori_url string) {

	// 获取书籍索引信息
	selection := get_html_content(ori_url, "body .wrap #pleft")
	selection.Find("dl[id=plist]").EachWithBreak(func(i int, s *goquery.Selection) bool {
		name := s.Find("dt a").Text()
		url := s.Find("dt a").AttrOr("href", "")
		fmt.Printf("sort %d: %s - %s\n", i, name, url)
		zxcs_get_download(url)
		return true
	})

	flag := selection.Find("div[id=pagenavi] span").Text()
	intFlag, err := strconv.Atoi(flag)
	if err != nil {
		log.Fatal(err)
	}

	strFlag := strconv.Itoa(intFlag + 1)

	selection.Find("div[id=pagenavi] a:contains('" + strFlag + "')").EachWithBreak(func(i int, s *goquery.Selection) bool {
		name := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("pagenavi %d: %s - %s\n", i, name, url)
		zxcs_get_sort(url)
		return false
	})
}

func zxcs8_get_latest() {

	// 获取最新的书籍信息
	get_html_content(ROOT_URL_ZXCS8, ".wrap #content a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		subUrl := s.AttrOr("href", "")
		// title := s.Find("i").Text()
		fmt.Printf("latest %d: %s - %s\n", i, band, subUrl)
		zxcs_get_download(subUrl)
	})
}

func zxcs_get_download(ori_url string) { //down_2

	//获取书籍下载页面
	get_html_content(ori_url, "body .wrap #pleft .pagefujian .down_2 a").Each(func(i int, s *goquery.Selection) {
		// book_downurl := s.Text()
		// fmt.Printf("download: %s \n", s.AttrOr("href", ""))
		zxcs_get_rar(s.AttrOr("href", ""))
	})
}

func zxcs_get_rar(ori_url string) { //downfile

	//获取书籍下载链家
	get_html_content(ori_url, "body .wrap .content .downfile  a").Each(func(i int, s *goquery.Selection) {
		line := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("line: %s  url: %s\n", line, url)
		write_file(url, OUT_FILE+line+".txt")
	})
}
