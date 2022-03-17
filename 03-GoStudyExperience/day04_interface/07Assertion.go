package main

import "fmt"

type Point struct {
	x int
	y int
}

func main07() {
	var i interface{}
	var point = Point{
		x: 19,
		y: 39,
	}
	i = point
	j := i.(Point) //类型断言
	fmt.Printf("j type is %T value is %v\n", j, j)

	var t = 100
	i = t
	k := i.(int)
	fmt.Println("k = ", k)

	var f = 100.0
	i = f
	temp, ok := i.(float32)
	if !ok {
		fmt.Println("err.....")
	}
	fmt.Println("temp = ", temp)

}
