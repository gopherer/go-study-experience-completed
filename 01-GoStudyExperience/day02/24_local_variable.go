package main 

import "fmt"

func test01(){
	a := 19
	fmt.Println("a = ",a)
}
func main24(){
	test01()
	//定义在{}里面的变量就是局部变量，只能在{}里面生效
	//执行到定义变量的那句话，才开始分配空间，离开作用域自动释放空间
	//作用域，变量其作用的范围
	//a = 99 error 
	{
		i := 19
		fmt.Println("i = ",i)
	}
	//i = 88 error
	if flag := 9; flag ==9 {
		fmt.Println("flag = ",flag)
	}
	//flag = 99//error
}