// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// //设置随机数种子 
// func Init(s []int) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < len(s); i++ {
// 		s[i] = rand.Intn(100)
// 	}
// }

// //排序.
// func BobbleSort(s []int) {
// 	n := len(s)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-1-i; j++ {
// 			if s[j] > s[j+1] {
// 				s[j], s[j+1] = s[j+1], s[j]
// 			}
// 		}
// 	}
// }
// func main22() {
// 	n := 10
// 	s := make([]int, n)

// 	Init(s)
// 	fmt.Println("排序前 ", s)
// 	BobbleSort(s)
// 	fmt.Println("排序后 ", s)
// }
