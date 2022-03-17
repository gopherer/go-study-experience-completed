package main

//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//type Process struct {
//	name     byte
//	cupTime  int
//	needTime int
//	priority int
//	state    string
//}
//
//const processCount = 5
//
//var pQueue [processCount]Process
//
//func main() {
//	var arr [processCount]int
//	var chQueue = make([]chan int, processCount)
//	var needTimeCount int
//
//	rand.Seed(time.Now().UnixNano())
//
//	for i := 0; i < processCount; i++ {
//		chQueue[i] = make(chan int)
//	}
//	for i := 0; i < processCount; i++ {
//		go pQueue[i].HandelProcess(i, chQueue[i])
//	}
//
//	for i := 0; i < processCount; i++ {
//		arr[i] = pQueue[i].priority
//		needTimeCount += pQueue[i].needTime
//	}
//	UI()
//	for i := 0; i < needTimeCount; i++ {
//		t := argMax(arr)
//		for j := 0; j < processCount; j++ {
//			if t == j {
//				chQueue[j] <- t
//				break
//			}
//		}
//		for i := 0; i < processCount; i++ {
//			arr[i] = pQueue[i].priority
//		}
//	}
//}
//
//func (process *Process) HandelProcess(i int, ch <-chan int) {
//	process.name = byte('a' + i)
//	process.cupTime = 0
//	process.needTime = rand.Intn(3) + 1
//	process.priority = rand.Intn(10) + 40
//	process.state = "waiting"
//	for {
//		select {
//		case <-ch:
//			process.priority -= 3
//			process.needTime -= 1
//			process.cupTime += 1
//			if process.needTime != 0 {
//				process.state = "running"
//			} else {
//				process.state = "terminated"
//				process.priority = -1
//				UI()
//				goto flag
//			}
//			UI()
//		}
//	}
//flag:
//	for {
//		<-ch
//	}
//}
//
//func UI() {
//	fmt.Println("name  cupTime  needTime  priority  state")
//	for _, v := range pQueue {
//		fmt.Printf("%c", v.name)
//		fmt.Println("      ", v.cupTime, "      ", v.needTime, "      ", v.priority, "      ", v.state)
//	}
//	fmt.Println("-----------------------------------------------------------------")
//}
//
//func argMax(arr [processCount]int) int {
//	maxVal := arr[0]
//	maxIndex := 0
//	for i := 1; i < len(arr); i++ {
//		if maxVal < arr[i] {
//			maxVal = arr[i]
//			maxIndex = i
//		}
//	}
//	return maxIndex
//}
