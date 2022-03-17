package main

import "fmt"

func main11() {
	var a int
	fmt.Printf("请输入变量a: ")

	//阻塞等待用户输入
	//fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Println("a = ", a)

}
