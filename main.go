// test
package main

import (
	"fmt"
	"spider"
)

func test_func(param map[string]string) {
	fmt.Println(param["key"])
}

func main() {
	fmt.Println("==================================开始======================================")

	// spider.Zxcs8()
	// spider.Fangxie()
	// spider.Fuli()

	param := make(map[string]string)
	param["key"] = "hello world!!"
	spider.Start_work(test_func, param, 10)

	fmt.Println("==================================结束======================================")
}
