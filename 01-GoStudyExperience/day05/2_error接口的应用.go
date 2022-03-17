package main

import (
	"errors"
	"fmt"
)

//MyDiv 两数相除
func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}
func main02() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}
}
