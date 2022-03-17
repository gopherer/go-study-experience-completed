package main

import "fmt"

func main13() {

	//给int64起一个别名叫bigint
	type bigint int64
	var a bigint
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)
	var b long
	var ch char = 'a'
	fmt.Printf("b = %d ch = %c\n", b, ch)
}
