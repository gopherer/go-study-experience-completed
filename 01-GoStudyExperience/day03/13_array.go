package main

import "fmt"

//数组做函数参数，是值传递
//实参数值的每一个元素都会拷贝一份过去
//形参的数字是实参的数组的复制品
func array(a [5]int) {
	a[0] = 888
	fmt.Println("array a = ", a)
}
func main13() {
	a := [5]int{1, 2, 3, 4, 5}
	array(a)
	fmt.Println("main a = ", a)
	array1(&a)
	fmt.Println("main a = ", a)
}

//*p指向数组a，为数组指针
//*p就代表实参a
func array1(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("array1 *p ", *p)
}
