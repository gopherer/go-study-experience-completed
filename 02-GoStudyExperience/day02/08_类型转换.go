package main

import "fmt"

func main08() {
	//数据类型(变量)
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)
	sum := a + b + c
	fmt.Printf("sum = %.2f\n", float64(sum)/3)
}
func main0801() {
	var a float64 = 1.22
	var b float32 = 22.33
	//建议低精度转高精度 可以保证数据不丢失
	fmt.Printf("a+b = %.2f\n", a+float64(b))
	fmt.Println(a + float64(b))

	var c float64 = 3.999
	fmt.Println(int(c)) //3 保留整数位，丢弃小数位
}
func main0802() {
	//1057453秒 几天几小时几分钟几秒
	var a int = 1057453
	fmt.Println("天:", a/60/60/24%365)
	fmt.Println("时:", a/60/60%24)
	fmt.Println("分:", a/60%60)
	fmt.Println("秒:", a%60)
	sum := 12*24*60*60 + 5*60*60 + 44*60 + 13
	if sum == a {
		fmt.Println(true)
	}
}
