package main

import "fmt"

func main02() {
	var p *int
	fmt.Println("p = ", p) //默认值为nil

	///*p = 666 error ，因为p没有合法的指向

	var a int
	p = &a
	*p = 888
	fmt.Println("a = ", a)
}
