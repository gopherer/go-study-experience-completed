package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Process struct {
	name     byte
	cupTime  int
	needTime int
	round    int
	count    int
	state    string
}

const processCount = 5

var pro = make([]Process, processCount)

func main() {
	var needTimeCount = 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < processCount; i++ {
	Flag:
		temp := rand.Intn(6) + 1
		if temp%2 == 0 {
			pro[i].needTime = temp
		} else {
			goto Flag
		}
		pro[i].name = byte(97 + i)
		pro[i].cupTime = 0
		pro[i].round = 0
		pro[i].count = 0
		pro[i].state = "waiting"
	}
	for i := 0; i < processCount; i++ {
		needTimeCount += pro[i].needTime
	}
	for i := 0; i < needTimeCount/2; i++ {
		t := i % processCount
		if pro[t].needTime != 0 {
			pro[t].HandleProcess()
		} else {
			for i := 0; i < processCount; i++ {
				t = (t + 1) % processCount
				if pro[t].needTime != 0 {
					pro[t].HandleProcess()
					break
				}
			}
		}
	}
}

func (pro *Process) HandleProcess() {
	pro.needTime -= 2
	pro.cupTime += 2
	pro.round += 1
	pro.count += 1
	if pro.needTime != 0 {
		pro.state = "running"
		UI()
	} else {
		pro.state = "terminated"
		UI()
	}
}
func UI() {
	fmt.Println("name  cupTime  needTime  round  count  state")
	for i := 0; i < processCount; i++ {
		fmt.Printf("%c         %d     %d        %d       %d     %s\n", pro[i].name, pro[i].cupTime, pro[i].needTime, pro[i].round, pro[i].count, pro[i].state)
	}
	fmt.Println("-----------------------------------------------------")
}
