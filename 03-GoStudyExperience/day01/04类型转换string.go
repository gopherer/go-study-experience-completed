package main

import (
	"fmt"
	"strconv"
)

func main04() {
	var i int = 100
	var f float64 = .888
	var b bool = true
	var s string

	s = fmt.Sprintf("%d", i) //%v 亦可以
	fmt.Printf("i type is %T and i value is %s\n", s, s)

	s = fmt.Sprintf("%f", f)
	fmt.Printf("f type is %T and f value is %s\n", s, s)

	s = fmt.Sprintf("%t", b)
	fmt.Printf("b type is %T and b value is %v", s, s)
}
func main0401() {
	var i int = 100
	var f float64 = 9.88
	var b bool = true
	var s string

	s = strconv.FormatInt(int64(i), 10)
	fmt.Printf("i type is %T and i value is %s\n", s, s)

	s = strconv.FormatFloat(f, 'f', 3, 64)
	fmt.Printf("f type is %T and f value is %s\n", s, s)

	s = strconv.FormatBool(b)
	fmt.Printf("b type is %T and b value is %v\n", s, s)
}
