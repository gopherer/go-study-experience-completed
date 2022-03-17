package main

import "fmt"

//定义接口
type inter interface {
	CSocketProtocol() //通讯接口
	CEncDesProtocol() //加密接口
}

//厂商1
type CSckImp1 struct {
	socket string
	data   string
}

//厂商2
type CSckImp2 struct {
	socket string
	data   string
	value  int
}

func (cs1 *CSckImp1) CSocketProtocol() {
	fmt.Printf("厂商1的通讯接口数据为:%s\n", cs1.socket)
}
func (cs1 *CSckImp1) CEncDesProtocol() {
	fmt.Printf("厂商1的加密接口数据为:%s\n", cs1.data)
}
func (cs2 *CSckImp2) CSocketProtocol() {
	fmt.Printf("厂商2的通讯接口数据为:%s\n", cs2.socket)
}
func (cs2 *CSckImp2) CEncDesProtocol() {
	fmt.Printf("厂商2的加密接口数据为:%s 价值为:%d\n", cs2.data, cs2.value)
}

//实现多态
func framework(i inter) {
	i.CSocketProtocol()
	i.CEncDesProtocol()
}
func main04() {
	cs1 := CSckImp1{"厂商1的通讯数据", "厂商1的加密数据"}
	framework(&cs1)
	cs2 := CSckImp2{"厂商2的通讯数据", "厂商2的加密数据", 100}
	framework(&cs2)
}
