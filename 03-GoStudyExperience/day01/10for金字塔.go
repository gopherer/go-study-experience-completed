package main

import "fmt"

func main10() {
	//由简到难
	//打印一个矩形
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//打印半个金字塔
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//打印金字塔
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//打印空心金字塔
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || k == 3 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	//变活
	var level int = 10
	for i := 1; i <= level; i++ {
		for j := 1; j <= level-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == level {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
