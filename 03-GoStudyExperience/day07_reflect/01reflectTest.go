package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	AGe  int
}

func main01() {
	//num := 100
	//refTest(num)
	//fmt.Println(num)
	stu := Student{
		Name: "jack",
		AGe:  18,
	}
	refTest02(stu)

}
func refTest(b interface{}) {
	//通过反射获取传入变量的type
	rTpye := reflect.TypeOf(b)
	fmt.Printf("rTpye Type %T rTpye value %v\n", rTpye, rTpye)
	kind := rTpye.Kind()
	fmt.Println(kind)
	//获取reflect.Value
	refValue := reflect.ValueOf(b)
	fmt.Printf("refValue Type %T refValue value %v\n", refValue, refValue)
	kind = refValue.Kind()
	fmt.Println(kind)
	n := 10
	num := refValue.Int()
	fmt.Printf("num Type %T num value %v\n", num, num)
	fmt.Println(int64(n) + num)

	iv := refValue.Interface()
	fmt.Printf("iv Type %T iv value %v\n", iv, iv)
	num2 := iv.(int)
	fmt.Printf("num2 Type %T num2 value %v\n", num2, num2)
}
func refTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("rTpye Type %T rTpye value %v\n", rType, rType)
	kind := rType.Kind()
	fmt.Println(kind)

	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue Type %T rValue value %v\n", rValue, rValue)
	kind = rValue.Kind()
	fmt.Println(kind)
	iv := rValue.Interface()
	fmt.Printf("iv type = %T value = %v\n", iv, iv)
	t, ok := iv.(Student)
	if ok {
		fmt.Println(t.Name)
	}
}
