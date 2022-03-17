package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:/GoCode/src/03_gocode/day05_file/test.txt")
	if err != nil {
		fmt.Println("文件打开失败..")
		return
	}
	defer func() {
		err = file.Close()
		fmt.Println("文件关闭成功")
	}()
	//创建缓冲区
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("文件读取失败")
		}
		fmt.Print(str)
	}
}
