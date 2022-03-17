package main

import (
	"errors"
	"fmt"
)

type circleQueue struct {
	maxSize int
	array   [4]int
	head    int //0
	tail    int //0
}

func (cQ *circleQueue) push(value int) (err error) {
	if cQ.isFull() {
		return errors.New("circleQueue is full ")
	}
	cQ.array[cQ.tail] = value
	cQ.tail = (cQ.tail + 1) % cQ.maxSize
	return nil
}
func (cQ *circleQueue) pop() (value int, err error) {
	if cQ.isEmpty() {
		err = errors.New("queue is empty")
		return 0, err
	}
	value = cQ.array[cQ.head]
	cQ.head = (cQ.head + 1) % cQ.maxSize
	return
}
func (cQ *circleQueue) listQueue() {
	if cQ.size() == 0 {
		fmt.Println("queue is empty")
		return
	}
	fmt.Println("queue detail:")
	temp := cQ.head
	for i := 0; i < cQ.size(); i++ {
		fmt.Printf("array[%d] = %d\n", temp, cQ.array[temp])
		temp = temp + 1%cQ.maxSize
	}
}
func (cQ *circleQueue) size() int {
	return (cQ.tail - cQ.head + cQ.maxSize) % cQ.maxSize
}
func (cQ *circleQueue) isFull() bool {
	return (cQ.tail+1)%cQ.maxSize == cQ.head
}
func (cQ *circleQueue) isEmpty() bool {
	return cQ.tail == cQ.head
}
func main() {
	cqueue := &circleQueue{
		maxSize: 4,
		array:   [4]int{},
		head:    0,
		tail:    0,
	}
	var key int
	for {
		fmt.Println("1 pushQueue 2popQueue 3listQueue")
		_, err := fmt.Scanln(&key)
		if err != nil {
			return
		}
		switch key {
		case 1:
			var t int
			fmt.Println("请输入要添加的变量：")
			_, err := fmt.Scanln(&t)
			err = cqueue.push(t)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			value, err := cqueue.pop()
			fmt.Println(value)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			cqueue.listQueue()
		}

	}
}
