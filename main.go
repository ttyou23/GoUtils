// test
package main

import (
	"fileutils"
	"fmt"
	"spider"
)

func test_func(param map[string]string) {
	fmt.Println(param["key"])
}

func main() {
	fmt.Println("==================================开始======================================")

	var cmd string

	// param := make(map[string]string)
	// param["key"] = "hello world!!"
	// spider.Start_work(test_func, param, 1)

	for true {
		fmt.Println("\n=====================主菜单==========================")
		fmt.Println("请输入命令：0(退出程序) 1(爬虫) 2(文件处理)")

		fmt.Scanln(&cmd)
		if cmd == "0" {
			break
		} else if cmd == "1" {
			fun_spider()
		} else if cmd == "2" {
			fun_file_format()
		}

	}

	fmt.Println("==================================结束======================================")

}

func fun_spider() {

	var cmd string
	for true {
		fmt.Println("\n=====================爬虫==========================")
		fmt.Println("请输入爬虫命令：0(回到主菜单) 1(知轩藏书) 2(福利小说) 3(房协网) 4(PV190)")
		fmt.Scanln(&cmd)

		if cmd == "0" {
			break
		} else if cmd == "1" {
			spider.Zxcs8()
		} else if cmd == "2" {
			spider.Fuli()
		} else if cmd == "3" {
			spider.Fangxie()
		} else if cmd == "4" {
			spider.PV190()
		}
	}

}

func fun_file_format() {

	var cmd, path string
	for true {
		fmt.Println("\n====================文件处理=========================")
		fmt.Println("请输入文件处理命令：0(回到主菜单) 1(加密) 2(解密)")
		fmt.Scanln(&cmd)
		if cmd == "0" {
			return
		}

		fmt.Println("请输入待处理的文件路径：备注（0：代码当前目录）")
		fmt.Scanln(&path)

		if cmd == "1" {
			fileutils.FileFormat(path, true)
		} else if cmd == "2" {
			fileutils.FileFormat(path, false)
		}
	}
}
