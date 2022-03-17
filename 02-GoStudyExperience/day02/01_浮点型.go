package main

import "fmt"

func main01() {
	//浮点型数据分为 单精度float32(精度 小数位7位) float64(精度 小数位15位)
	var a float32 = 3.15
	fmt.Println(a)
	fmt.Printf("a = %.2f\n", a)

	b := 4.24
	fmt.Println(b)
	fmt.Printf("b = %T\n", b) //类型位float64
}

//买西瓜
func main0101() {
	price := 3.5
	var num float64
	fmt.Scan(&num)
	sum := price * num
	fmt.Println(sum)
}
