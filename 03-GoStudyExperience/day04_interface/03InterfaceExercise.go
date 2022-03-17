package main

import "fmt"

type Usb interface {
	Say()
}
type Student struct {
}

func (s *Student) Say() {
	fmt.Println("Say()")
}
func main() {
	var stu = Student{}
	//var usb Usb = stu //error say()方法的接收者为&stu
	var usb Usb = &stu
	usb.Say()
}
