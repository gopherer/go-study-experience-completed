package main

import "fmt"

func main07() {
	var a int
	fmt.Println(&a)

	fmt.Scan(&a)
	fmt.Printf("a = %d\n", a)

}
func main0701() {
	var a, b int

	fmt.Scan(&a, &b) //空格和回车 可以作为一个数据输入完成的标志
	fmt.Println(a, b)
}
func main0702() {
	var a int
	var b string
	//格式化输入
	fmt.Scanf("%d %s", &a, &b) //空格作为一个数据输入完成的标志，因为回车会被b接收
	fmt.Printf("a = %d b = %s", a, b)

}
