package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("D:\\新建文件夹\\6.9更新总库.txt")
	if err != nil {
		fmt.Print("open file err")
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("file close err")
			os.Exit(1)
		}
	}()
	reader := bufio.NewReader(file)
	tem := ""
	fmt.Println("输入qq号或电话号")
	_, err = fmt.Scanln(&tem)
	if err != nil {
		fmt.Println("输入有误.")
		return
	}
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("file end")
			return
		}
		flag := strings.Contains(str, tem)
		if flag {
			fmt.Print("qq号---电话号\n", str)
			continue
		}
	}
}
