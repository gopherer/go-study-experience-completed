package main

import "fmt"

//厂商2
type CSckImp2 struct {
	socket string
	data   string
	value  int
}

func (cs2 *CSckImp2) CSocketProtocol() {
	fmt.Printf("厂商2的通讯接口数据为:%s\n", cs2.socket)
}
func (cs2 *CSckImp2) CEncDesProtocol() {
	fmt.Printf("厂商2的加密接口数据为:%s 价值为:%d\n", cs2.data, cs2.value)
}
