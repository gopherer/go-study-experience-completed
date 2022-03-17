package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	maxTop int
	top    int
	arr    [5]int
}

func (stack *Stack) push(val int) (err error) {
	if stack.maxTop-1 == stack.top {
		return errors.New("this stack is full")
	}
	stack.top++
	stack.arr[stack.top] = val
	return
}
func (stack *Stack) pop() (val int, err error) {
	if stack.top == -1 {
		return 0, errors.New("stack is empty")
	}
	val = stack.arr[stack.top]
	stack.top--
	return
}
func (stack *Stack) list() (err error) {
	if stack.top == -1 {
		return errors.New("stack is empty")
	}
	for i := stack.top; i > -1; i-- {
		fmt.Printf("%d ---> %d\n", i, stack.arr[i])
	}
	return
}
func main() {
	stack := &Stack{
		maxTop: 5,
		top:    -1,
	}
	for i := 0; i < 5; i++ {
		err := stack.push(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	err := stack.list()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	for i := 0; i < 3; i++ {
		_, err := stack.pop()
		if err != nil {
			fmt.Println(err)
		}
	}
	err = stack.list()
	if err != nil {
		fmt.Println(err)
	}

}
