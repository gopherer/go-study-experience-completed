package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main04() {
	var File1path = "C:/GoCode/src/03_gocode/day05_file/test.txt"
	var File2path = "C:/GoCode/src/03_gocode/day05_file/text1.txt"
	data, err := ioutil.ReadFile(File1path)
	if err != nil {
		fmt.Println("文件读取失败。")
		return
	}
	err = ioutil.WriteFile(File2path, data, 0666)
	if err != nil {
		fmt.Println("文件写入失败")
	}

	b, err := PathExists(File1path)
	if b == true {
		fmt.Println("文件存在")
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil //文件或者目录存在
	}
	if os.IsNotExist(err) {
		return false, nil //文件或目录不存在
	}
	return false, nil //无法判断文件是否存在
}
