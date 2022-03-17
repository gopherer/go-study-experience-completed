package main

import "fmt"

type Usb2 interface {
	Start()
	Stop()
}
type Computer1 struct {
}
type Phone1 struct {
	Name string
}
type Camera1 struct {
	Name string
}

func (c Computer1) Working(usb Usb2) {
	usb.Start()
	if Phone1, ok := usb.(Phone1); ok {
		Phone1.Call()
	}
	usb.Stop()
}
func (p Phone1) Start() {
	fmt.Println("Phone1 开始工作....")
}
func (p Phone1) Stop() {
	fmt.Println("Phone1 停止工作...")
}
func (p Phone1) Call() {
	fmt.Println("Phone1 在打电话")
}
func (c Camera1) Start() {
	fmt.Println("Camera1 开始工作....")
}
func (c Camera1) Stop() {
	fmt.Println("Camera1 停止工作....")
}
func main08() {
	//定义一个Usb接口数组，可以存放Phone1和Camera1结构体变量
	//这里就体现出多态数组
	var UsbArr [3]Usb2
	UsbArr[0] = Phone1{Name: "iPhone1"}
	UsbArr[1] = Phone1{Name: "vivo"}
	UsbArr[2] = Camera1{Name: "尼康"}
	fmt.Println(UsbArr)
	UsbArr[1].Start()
	UsbArr[1].Stop()
	fmt.Println("---------------")
	var computer = Computer1{}
	for _, value := range UsbArr {
		computer.Working(value)
	}

}
