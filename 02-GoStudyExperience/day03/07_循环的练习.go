package main

import "fmt"

func main07() {
	//敲7
	for i := 1; i <= 100; i++ {
		if i%7 == 0 || i%10 == 7 || i/7 == 7 {
			fmt.Println("敲桌子")
		} else {
			fmt.Println(i)
		}
	}
}
func main0701() {
	//水仙花数
	for i := 100; i <= 999; i++ {
		//百位
		a := i / 100
		//十位
		b := i / 10 % 10
		//个位
		c := i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}

//三角形
func main0702() {
	var l int = 10
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
