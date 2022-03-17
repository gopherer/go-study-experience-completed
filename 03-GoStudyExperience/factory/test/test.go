package test

type student struct {
	Name string
	age  int
}
//利用NewStudent函数来实现调用私有结构体   newStudent 首字母为小写 私有 无法为其他包提供服务
func NewStudent(name string) *student {
	return &student{
		Name:name,
	}
}
//GetAge 首字母大写 可以为其他包所使用 为student结构体专用方法 用此来解决结构体字段首字母小写（私有）问题
func (stu *student)GetAge(age int)int{
	stu.age = age
	return stu.age
}
