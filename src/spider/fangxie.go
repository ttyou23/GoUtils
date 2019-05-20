// fangxie
package spider

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ROOT_URL_FANGXIE string = "https://www.cdfangxie.com/"
const FANGXIE_HOUSE_EXCEL = "fangxie.xlsx"
const FANGXIE_HOUSE_URL = "fangxie_url.txt"

var houses [][]string
var HOUSE_TITLE []string = []string{"区域", "项目名称", "项目咨询电话", "预/现售证号", "房屋用途", "预售面积(平方米)", "上市时间", "购房登记规则下载地址", "成品住房装修方案价格表下载地址", "项目网址"}

func Fangxie() {
	houses = append(houses, HOUSE_TITLE)
	fangxie_tag()
	save2excel(houses, FANGXIE_HOUSE_EXCEL)
}

func fangxie_tag() {
	get_html_content(ROOT_URL_FANGXIE, "body div[class=cont1_rukou] a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := s.Text()
		url := s.AttrOr("href", "")
		fmt.Printf("title: %s - %s\n", text, url)
		fangxie_sort(ROOT_URL_FANGXIE + url)
		return false
	})
}

func fangxie_sort(ori_url string) {
	selection := get_html_content(ori_url, "body .wrapper .main .right_cont")
	selection.Find(".ul_list li a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title := s.Text()
		url := ROOT_URL_FANGXIE + s.AttrOr("href", "")
		zone := title[:strings.Index(title, "|")]
		fmt.Printf("fangxie_sort: %s - %s\n", zone, url)
		fangxie_info(url, zone)
		return false
	})

	selection.Find(".pages2 a:contains('下一页')").Each(func(i int, s *goquery.Selection) {
		url := ROOT_URL_FANGXIE + s.AttrOr("href", "")
		fmt.Printf("fangxie_sort: %s - %s\n", s.Text(), url)
		// fangxie_sort(url)
	})
}

func fangxie_info(ori_url, zone string) {

	item := make([]string, 0)
	item = append(item, zone)

	get_html_content(ori_url, "body .main .infor p").Each(func(i int, s *goquery.Selection) {
		text := strings.ReplaceAll(strings.ReplaceAll(s.Text(), "\n", ""), " ", "")
		// fmt.Printf("fangxie_info: %s len: %d\n", text, len(text))
		if strings.Index(text, "项目名称") > 0 {
			item = append(item, text[strings.Index(text, ":")+1:])
			fmt.Printf("%s\n", text)
		} else if strings.Index(text, "项目咨询电话") > 0 {
			item = append(item, text[strings.Index(text, ":")+1:])
		} else if strings.Index(text, "预/现售证号") > 0 {
			item = append(item, text[strings.Index(text, ":")+1:])
		} else if strings.Index(text, "房屋用途") > 0 {
			item = append(item, text[strings.Index(text, ":")+1:])
		} else if strings.Index(text, "预售面积(平方米)") > 0 {
			item = append(item, text[strings.Index(text, ":")+1:])
		} else if strings.Index(text, "上市时间") > 0 {
			item = append(item, text[strings.Index(text, ":")+1:])
		} else if strings.Index(text, "购房登记规则") > 0 {
			url := s.Find("a").AttrOr("href", "")
			item = append(item, url)
			write_file(url, FANGXIE_HOUSE_URL)
		} else if strings.Index(text, "成品住房装修方案价格表点击下载") > 0 {
			url := s.Find("a").AttrOr("href", "")
			item = append(item, url)
			write_file(url, FANGXIE_HOUSE_URL)
		}
	})

	if len(item) < 10 {
		item = append(item, "")
	}
	item = append(item, ori_url)
	houses = append(houses, item)
}
