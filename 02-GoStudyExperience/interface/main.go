package main

//定义接口
type inter interface {
	CSocketProtocol() //通讯接口
	CEncDesProtocol() //加密接口
}

//实现多态
func framework(i inter) {
	i.CSocketProtocol()
	i.CEncDesProtocol()
}
func main() {
	cs1 := CSckImp1{"厂商1的通讯数据", "厂商1的加密数据"}
	framework(&cs1)
	cs2 := CSckImp2{"厂商2的通讯数据", "厂商2的加密数据", 100}
	framework(&cs2)
}
