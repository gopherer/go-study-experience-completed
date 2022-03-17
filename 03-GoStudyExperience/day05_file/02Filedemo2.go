package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile来一次性读取文件 适用于小文件
	var filePath = "C:/GoCode/src/03_gocode/day05_file/test.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err = ", err)
		return
	}
	fmt.Print(string(content))
}
