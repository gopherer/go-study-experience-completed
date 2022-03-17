package main

import "fmt"

func main17() {
	//for 初始条件; 判断条件; 条件变化{
	//}
	//1+2+3 ...100
	sum := 0
	//1)初始化条件 i :=
	//2)判断条件是否为真，i >= 100 ,如果为真执行循环体，如果为假，跳出循环体
	//3)条件变化 i++
	//4) 重复 2，3，4
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}
