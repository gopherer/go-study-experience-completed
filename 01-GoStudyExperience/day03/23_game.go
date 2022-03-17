package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	*p = num
}
func GetNum(s []int, num int) {
	s[0] = num / 1000       //千
	s[1] = num % 1000 / 100 //百
	s[2] = num % 100 / 10   //十
	s[3] = num % 10         //个
}
func OnGame(randNum []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for { //用户输入数字
			fmt.Println("请输入一个4位数：")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输数字有误 请重新输入：")
		}

		GetNum(keySlice, num)
		//fmt.Println("keySlice = ", keySlice)
		n := 0
		for i := 0; i < 4; i++ { //判断是否与随机数相同
			if keySlice[i] > randNum[i] {
				fmt.Printf("第%d位大了\n", i+1)
			} else if keySlice[i] < randNum[i] {
				fmt.Printf("第%d位小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了了\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("你赢了")
			break
		}
	}

}
func main23() {
	var randNum int
	CreatNum(&randNum)
	//fmt.Println("randNum = ", randNum)
	randSlice := make([]int, 4)
	GetNum(randSlice, randNum)
	//fmt.Println("randSlice = ", randSlice)
	OnGame(randSlice)
}
