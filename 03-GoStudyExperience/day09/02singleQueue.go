package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (Q *Queue) addQueue(value int) (err error) {
	if Q.rear == Q.maxSize-1 {
		err = errors.New("queue is full")
		return
	}
	Q.array[Q.rear+1] = value
	Q.rear++
	return
}
func (Q *Queue) getQueue() (value int, err error) {
	if Q.rear == Q.front {
		err = errors.New("queue is empty")
		return -1, err
	}
	Q.front++
	value = Q.array[Q.front]
	return
}
func (Q *Queue) showQueue() {
	fmt.Println("queue detail:")
	for i := Q.front + 1; i <= Q.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, Q.array[i])
	}
}
func main() {
	queue := &Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
	}
	var key int
	for {
		fmt.Println("1 addQueue 2getQueue 3showQueue")
		_, err := fmt.Scanln(&key)
		if err != nil {
			return
		}
		switch key {
		case 1:
			var t int
			fmt.Println("请输入要添加的变量：")
			_, err := fmt.Scanln(&t)
			err = queue.addQueue(t)
			if err != nil {
				fmt.Println(err)
				return
			}
		case 2:
			value, err := queue.getQueue()
			fmt.Println(value)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			queue.showQueue()
		}

	}
}
