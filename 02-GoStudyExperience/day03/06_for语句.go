package main

import "fmt"

func main06() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello go")
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	//计算1-100偶素的和
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	sum = 0
	//计算1-100偶素的和
	for i := 0; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}
