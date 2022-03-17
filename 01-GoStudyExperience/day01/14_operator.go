package main

import "fmt"

func main14() {

	fmt.Println("4 > 3 结果：", 4 > 3)
	fmt.Println("4 != 3 结果：", 4 != 3)

	fmt.Println("!(4 > 3) 结果：", !(4 > 3))
	fmt.Println("!(4 != 3) 结果：", !(4 != 3))
	//&& 与， 并且，左右都为真，结果才为真，其余都为假
	//fmt.Println("true && true 结果：", true && true)
	fmt.Println("true && flase 结果：", true && false)
	//||,或者，左右都为假时，结果才为假，其余都为真
	fmt.Println("true || false 结果：", true || false)
	//fmt.Println("false || false 结果：", false || false)

	a := 8
	//fmt.println("0 <= a <= 10 结果：", 0 <= a <= 10)//err
	fmt.Println("0 <= a && a <= 10 结果：", 0 <= a && a <= 10)
	a += 2
	fmt.Println("a = ", a)
	a++
	fmt.Println("a = ", a)

}
