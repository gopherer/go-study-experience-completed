package main

import "fmt"
//实现1+2+3+...100
func function21() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func function22(i int) int {
	if i == 1 {
		return 1
	}
	return i + function22(i-1)
}

func function23(i int) int {
	if i == 100 {
		return 100
	}
	return i + function23(i+1)
}

func main2109() {
	var sum int
	sum = function21()
	fmt.Println("sum = ", sum)

	sum = function22(100)
	fmt.Println("sum = ", sum)

	sum = function23(1)
	fmt.Println("sum = ", sum)
}
