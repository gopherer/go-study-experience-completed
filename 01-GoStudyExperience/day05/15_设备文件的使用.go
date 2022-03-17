package main

import (
	"fmt"
	"os"
)

func main15() {
	// os.Stdout.Close()//关闭后，无法输出到屏幕
	//标准设备文件(os.Stdout),默认已经打开，用户可以直接使用
	fmt.Println("are u ok?") //往标准输出设备(屏幕)写内容
	//os.Stdout
	os.Stdout.WriteString("u are ok\n")

	//os.Stdin.Close()//关闭后无法输入
	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a) //从标准输入设备中读取内容，放入a中
	fmt.Println("a = ", a)

}
