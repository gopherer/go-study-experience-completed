package main

import (
	"errors"
	"fmt"
)

func main01() {
	var err1 error = fmt.Errorf("%s", "this is normal err")
	fmt.Println("err1 = ", err1)
	err2 := fmt.Errorf("%s", "this is a normal err2")
	fmt.Println("err2 = ", err2)

	err3 := errors.New("this is normal err3")
	fmt.Println("err3 = ", err3)
}
