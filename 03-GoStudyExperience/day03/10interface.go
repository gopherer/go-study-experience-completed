package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Computer struct {
}
type Phone struct {
}
type Camera struct {
}

func (c Computer) Working(usb Usb) {
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

func main10() {

	var computer = Computer{}
	var phone = Phone{}
	var camera = Camera{}

	computer.Working(phone)
	computer.Working(camera)

}
