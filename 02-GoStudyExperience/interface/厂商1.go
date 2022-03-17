package main

import "fmt"

//厂商1
type CSckImp1 struct {
	socket string
	data   string
}

func (cs1 *CSckImp1) CSocketProtocol() {
	fmt.Printf("厂商1的通讯接口数据为:%s\n", cs1.socket)
}
func (cs1 *CSckImp1) CEncDesProtocol() {
	fmt.Printf("厂商1的加密接口数据为:%s\n", cs1.data)
}
