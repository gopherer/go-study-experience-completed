package main

import "fmt"

type Usb1 interface {
	Start()
	Stop()
}
type Computer struct {
}
type Phone struct {
	Name string
}
type Camera struct {
	Name string
}

func (c Computer) Working(usb Usb1) {
	usb.Start()
	usb.Stop()
}
func (p Phone) Start() {
	fmt.Println("Phone 开始工作....")
}
func (p Phone) Stop() {
	fmt.Println("Phone 停止工作...")
}
func (c Camera) Start() {
	fmt.Println("Camera 开始工作....")
}
func (c Camera) Stop() {
	fmt.Println("Camera 停止工作....")
}
func main06() {
	//定义一个Usb接口数组，可以存放Phone和Camera结构体变量
	//这里就体现出多态数组
	var UsbArr [3]Usb1
	UsbArr[0] = Phone{Name: "iphone"}
	UsbArr[1] = Phone{Name: "vivo"}
	UsbArr[2] = Camera{Name: "尼康"}
	fmt.Println(UsbArr)
	UsbArr[1].Start()
	UsbArr[1].Stop()
}
