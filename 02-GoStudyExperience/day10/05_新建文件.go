package main

import (
	"fmt"
	"os"
)

func main05() {
	//os.Create(文件名) 文件名可以写绝对路径 也可以写相对路径
	//返回值 文件指针 错误信息
	fp, err := os.Create("a.txt") //./表示当前路径 ../表示上一级路径
	if err != nil {
		//文件创建失败 1、路径不存在 2、文件权限限制 3、程序打开文件上限 65535
		fmt.Println("err = ", err)
		return
	}
	//如果文件不关闭 会造成内存浪费
	defer fp.Close()

}
