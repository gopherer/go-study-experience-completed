package main

import (
	"fmt"
	"math"
)

func main0201(){
	var sun int = 10//等价于 var sun  = 10
	var a = 10
	fmt.Println(a)
	fmt.Println(sun)
	sun = sun+10
	fmt.Println(sun)
}

func main0202(){
	var sun int
	fmt.Println(sun) //未赋值 默认值为0
	sun = 100
	fmt.Println(sun)
}

func main0203(){
	var sun float64 = 2
	var value float64 = sun *sun *sun*sun*sun*sun*sun*sun*sun*sun
	fmt.Println(value)
	var val float64 = math.Pow(sun,10)
	fmt.Println(val)
}

