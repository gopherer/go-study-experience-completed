package main

import (
	"errors"
	"fmt"
)

func div(a int, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("runtime err 除数不能为0")
	} else {
		result = a / b
	}
	return
}
func main01() {
	a := 10
	b := 5
	value, err := div(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
