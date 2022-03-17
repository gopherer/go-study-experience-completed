package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main06() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main0601() {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(100)
	var value int
	for {
		fmt.Println("请输入数字")
		fmt.Scan(&value)

		if value > num {
			fmt.Println("输入数字过大")
		} else if value < num {
			fmt.Println("输入数字过小")
		} else {
			fmt.Println("恭喜猜对了")
			break
		}
	}
}
func main0602() {
	rand.Seed(time.Now().UnixNano())
	var red [6]int

	for i := 0; i < len(red); i++ {
		v := rand.Intn(33) + 1
		for j := 0; j < i; j++ {
			if v == red[j] {
				v = rand.Intn(33) + 1
				j = -1
			}
		}
		red[i] = v
	}
	fmt.Println("红球:", red, "蓝球:", rand.Intn(16)+1)
}
