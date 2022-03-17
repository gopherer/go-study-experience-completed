package main 

import"fmt"

func main2102(){
	function05()
	function06(1,5,6)
	function07(2,99994,986)
}
//...int类型 ...type不定参数类型
func function05(args ...int){//传递的实参可以是0个或者多个

}

func function06(args ...int){
	fmt.Println("len(args) = ", len(args))//获取用户传递参数的个数

	for i := 0; i < len(args); i++{
		fmt.Printf("args[%d] = %d ", i, args[i])
	}
	fmt.Println("")
	//返回2个值，第一个是下标，第二个是下标对应的数值
	for i ,data := range args{
		fmt.Printf("args[%d] = %d ",i ,data)
	}
	fmt.Println("")

	for i := range args{
		fmt.Printf("args[%d] = %d ",i ,args[i])
	}
	fmt.Println("")
}
//固定参数一定要传参数，不定参数根据需求传递
func function07(a int,args ...int){
	fmt.Println("a = ", a)

	for i := range args{
		fmt.Printf("args[%d] = %d ",i ,args[i])
	}
	fmt.Println("")

}
//不定参数只能作为形参做个的一个参数
// func function08(args ...int, a int){//error
// }
