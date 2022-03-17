package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filePath = "C:/GoCode/src/03_gocode/day05_file/text1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer func() {
		err = file.Close()
		fmt.Print("文件写入成功")
	}()
	//创建缓冲区
	writer := bufio.NewWriter(file)
	//将数据写入缓冲区
	var str = "hello go\n"
	for i := 0; i < 5; i++ {
		_, err = writer.WriteString(str)
		if err != nil {
			fmt.Println("文件写入失败.")
			return
		}
	}
	//Flush方法将缓冲中的数据写入下层的io.Writer接口。
	err = writer.Flush()
	if err != nil {
		fmt.Println("文件写入失败")
	}
}
