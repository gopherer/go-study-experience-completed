package main

import (
	"fmt"
	"io"
	"os"
)

func main17() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage:xxx srcFile desFile")
		return
	}

	srcFile := list[1]
	desFile := list[2]
	if srcFile == desFile {
		fmt.Println("源文件不能和目的文件同名")
		return
	}

	//只读方式打开源文件
	sF, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//新建目的文件
	dF, err1 := os.Create(desFile)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	//延迟操作 关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，往目的文件写，都多少写多少
	buf := make([]byte, 2*1024) //创建一个2k的缓冲区
	for {
		n, err := sF.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			if err == io.EOF {
				fmt.Println("已到文件结尾")
				break
			}
		}
		//往目的文件写，读多少写多少
		dF.Write(buf[:n])
	}

}
