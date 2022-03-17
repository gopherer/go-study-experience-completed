package main

import (
	"flag"
	"fmt"
)

func main08() {
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是在接收用户命令行中输入-u后面的参数
	//"u"就是-u指定参数
	//""，默认值
	//"用户名，默认为空"说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "p", 3306, "端口号，默认为3306")
	//必须调用flag.Parse()
	flag.Parse()
	fmt.Printf("user = %v,pwd = %v host = %v port = %v", user, pwd, host, port)
}
