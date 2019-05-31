// excel
package fileformat

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"spider"
	"strings"
)

func FormatExcel(path string) {
	fp, err := os.Stat(path)
	if err != nil {
		fmt.Println("Stat err:", err)
		return
	}
	if fp.IsDir() {

		// files, err := ListDir(path, ".TXT", true)
		files, err := ListDir(path, ".ext", true)
		fmt.Println("ListDir:", files, err)
		if err != nil {
			fmt.Println("ListDir err:", err)
			return
		}
		for _, filepath := range files {
			save_data(filepath.fullpath)
		}
	} else if strings.HasSuffix(strings.ToUpper(path), ".TXT") {
		save_data(path)
	}
}

func save_data(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var exceldata [][]string
	var EXCEL_TITLE []string = []string{"区站号", "年份", "月份", "降水量", "平均气温", "日照时数"}
	exceldata = append(exceldata, EXCEL_TITLE)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			if line == "" {
				break
			}
		}
		if line != "" {
			list_item := format_data(line)
			exceldata = append(exceldata, list_item)
		}
		fmt.Println(line)
	}
	spider.Save2excel(exceldata, strings.ReplaceAll(path, ".txt", ".xlsx"))
	fmt.Println("该文件转换成功：" + path)
}

func format_data(data string) []string {
	var list_item []string
	if data == "" {
		return list_item
	}

	list := strings.Split(data, " ")
	for i, item := range list {
		fmt.Println(i, item)
		list_item = append(list_item, item)
	}
	return list_item
}
