package main

import (
	"fmt"
	"time"
)

func main08() {
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			count++
			fmt.Println(i, j)
		}
	}
	fmt.Println(count)
}

//电子钟
func main0801() {
	//时
	for i := 0; i < 24; i++ {
		for j := 0; j < 60; j++ {
			for k := 0; k < 60; k++ {
				time.Sleep(time.Millisecond * 950)
				fmt.Printf("%d 时 %d 分 %d 秒\n", i, j, k)
			}
		}
	}
}
func main0802() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().YearDay())
}
func main0803() {
	//九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
