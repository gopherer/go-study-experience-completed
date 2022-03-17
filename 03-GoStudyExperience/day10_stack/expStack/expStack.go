package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	maxTop int
	top    int
	arr    [20]int
}

func main() {
	numStack := &Stack{
		maxTop: 20,
		top:    -1,
		arr:    [20]int{},
	}
	operStack := &Stack{
		maxTop: 20,
		top:    -1,
		arr:    [20]int{},
	}
	exp := "30+3*6-4"
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	keepNum := ""
	res := 0

	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0])
		if operStack.isOperation(temp) {
			if operStack.top == -1 {
				err := operStack.push(temp)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				if operStack.priority(operStack.arr[operStack.top]) >= operStack.priority(temp) {
					num1, _ = numStack.pop()
					num2, _ = numStack.pop()
					oper, _ = operStack.pop()
					res = numStack.cal(num1, num2, oper)
					err := numStack.push(res)
					if err != nil {
						fmt.Println(err)
					}
					err = operStack.push(temp)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					err := operStack.push(temp)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		} else {
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				err := numStack.push(int(val))
				if err != nil {
					fmt.Println(err)
				}
			} else if operStack.isOperation(int([]byte(exp[index+1 : index+2])[0])) {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				err := numStack.push(int(val))
				if err != nil {
					fmt.Println(err)
				}
				keepNum = ""
			}

		}
		if index == len(exp)-1 {
			break
		}
		index++
	}
	for {
		if operStack.top == -1 {
			break
		}
		num1, _ = numStack.pop()
		num2, _ = numStack.pop()
		oper, _ = operStack.pop()
		res = numStack.cal(num1, num2, oper)
		err := numStack.push(res)
		if err != nil {
			fmt.Println(err)
		}
	}
	res, _ = numStack.pop()
	fmt.Println(exp, "=", res)
}

func (stack *Stack) push(oper int) (err error) {
	if stack.maxTop-1 == stack.top {
		return errors.New("this stack is full")
	}
	stack.top++
	stack.arr[stack.top] = oper
	return
}
func (stack *Stack) pop() (oper int, err error) {
	if stack.top == -1 {
		return 0, errors.New("stack is empty")
	}
	oper = stack.arr[stack.top]
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
func (stack *Stack) isOperation(oper int) bool {
	if oper == 42 || oper == 47 || oper == 43 || oper == 45 {
		return true
	} else {
		return false
	}
}
func (stack *Stack) priority(oper int) (res int) {
	res = -1
	if oper == 47 || oper == 42 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return
}
func (stack *Stack) cal(num1 int, num2 int, oper int) int {
	switch oper {
	case 42:
		return num2 * num1
	case 47:
		return num2 / num1
	case 43:
		return num2 + num1
	case 45:
		return num2 - num1
	default:
		fmt.Println("operation err")
	}
	return -1
}
