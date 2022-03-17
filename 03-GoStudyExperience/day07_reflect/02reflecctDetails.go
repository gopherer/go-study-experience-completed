package main

import (
	"fmt"
	"reflect"
)

func main02() {
	var num int = 100
	reflect01(&num)
	fmt.Println(num)
}
func reflect01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("rType type %T value %v\n", rType, rType)
	fmt.Println(rType.Kind())
	t := rType.Elem()
	fmt.Printf("t Type %T value %v\n", t, t)

	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue type %T rValue %v\n", rValue, rValue)
	fmt.Println(rValue.Kind())
	rValue.Elem().SetInt(200)

}
