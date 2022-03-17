package main

import "fmt"

func main20() {
	//goto可以在任何地方，但不能跨函数使用
	fmt.Println("hhhhhhhhhhh")

	goto End //goto是关键字，End是用户起的名字

	//fmt.Println("jjjjjjjjjjj")
End:
	fmt.Println("ssssssssssss")

}
