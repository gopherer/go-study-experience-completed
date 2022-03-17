package main

import "02_gocode/packageuse/userinfo"

//import HHH "userinfo" 给包名userinfo起了一个别名 需使用HHH包名访问包 原包名失效 go默认从gopath下src文件夹下索引文件夹（包）

//通过命令编译程序
//go build 01_main.go 02_test.go 03_demo.go
//go run 01_main.go 02_test.go 03_demo.go
//go build ./src
//go run ./src
//在多文件编程时，在go build中选择配置 在配置选择目录指定到文件所在的目录级别
//需要到File中Setting下GOPATH设置相关路径 文件夹名须为src
//调用其他文件夹是需设置上一级目录
func main() {
	test(10, 20)
	demo()
	//如果调用其他文件夹的函数 需要包名.函数名  以及导入import"文件名"
	userinfo.UserCreat()
	userinfo.UserLogin()

}
