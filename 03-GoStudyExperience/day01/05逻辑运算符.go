package main

import "fmt"

func test() bool {
	fmt.Println("test......")
	return true
}
func main05() {
	var i int = 10
	//因为 9 < i 恒为true 对于||整体为真 test()不执行
	if 9 < i || test() {
		fmt.Println("OK")
	}
	//因为 9 > i 恒为false 对于&&整体为假 test()不执行
	if 9 > i && test() {
		fmt.Println("OK")
	}

}
