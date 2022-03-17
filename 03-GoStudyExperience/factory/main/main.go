package main

import (
	"03_gocode/factory/test"
	"fmt"
)
func main() {
	var  stu = test.NewStudent("ton")
	fmt.Println(stu.Name)

	var StuAge = stu.GetAge(50)
	fmt.Println(StuAge)
}
