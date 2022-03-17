package main

import (
	"fmt"
)

func main() {
	var myMap [8][7]int

	//1-->表示墙
	//0--》表示未走过的点
	//2-->终点
	//3--》表示通路
	//设置最上 最下面的
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	//myMap[1][2] = 1
	//myMap[2][2] = 1
	b := findMap(1, 1, &myMap)
	if b {
		fmt.Println("find it")
	} else {
		fmt.Println("don't find it")
	}
	for _, v := range myMap {
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
}
func findMap(i int, j int, myMap *[8][7]int) bool {
	//分析出什么情况下找到出口
	//设myMap[6][5] = 2
	if myMap[6][5] == 2 {
		return true
	} else {
		//说明要继续寻找
		if myMap[i][j] == 0 { //说明该点可以勘测
			//假设该点可通
			myMap[i][j] = 2
			if findMap(i+1, j, myMap) { //下
				return true
			} else if findMap(i, j+1, myMap) { //右
				return true
			} else if findMap(i-1, j, myMap) { //上
				return true
			} else if findMap(i, j-1, myMap) { //左
				return true
			} else {
				myMap[i][j] = 3
				return false
			}
		} else { //说明是墙或者死路
			return false
		}
	}
}
