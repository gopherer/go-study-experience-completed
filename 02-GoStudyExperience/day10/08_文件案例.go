package main

import (
	"io"
	"os"
)

func main08() {
	fp, err := os.Open("a.txt")
	if err != nil {
		return
	}
	defer fp.Close()
	fp1, err := os.Create("../a.txt")
	if err != nil {
		return
	}
	defer fp1.Close()
	//拷贝文件 通过read读取 write写入
	buf := make([]byte, 1024)
	for {
		n, err := fp.Read(buf)
		if err == io.EOF {
			break
		}
		fp1.Write(buf[:n])
	}

}
