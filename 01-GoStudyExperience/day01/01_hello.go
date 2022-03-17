package main // 代码包声明语句。

import "fmt" //系统包用来输出的
//入口函数
func main01() { //左括号必须与函数同行
	// 打印函数调用语句。用于打印输出信息。
	fmt.Println("hello world") //Println 会自动换行 调用函数大部分需要导入包

}

//1) go语言以包作为管理单位
//2) 每个文件必须要声明包
//3) 程序必要要有一个main包
//4) go语言结尾没有分号
