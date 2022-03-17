package main 

import "fmt"
//Boy 内容
type Boy struct {
	name string
	id int
}
func main17(){
	i := make([]interface{},3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Boy{"yoyo",666}
	//类型查询，类型断言
	//第一个返回值下标，第二个返回值下标对应的值，
	for index,data := range i{
		//第一个返回的是值，第二个返回判断结果的真假
		if value,ok := data.(int); ok == true{
			fmt.Printf("x[%d] 类型int 内容%d\n",index,value)
		}else if value,ok := data.(string); ok == true{
			fmt.Printf("x[%d] 类型string 内容%s\n",index,value)
		}else if value,ok := data.(Boy); ok == true{
			fmt.Printf("x[%d] 类型Boy 内容name = %s id = %d\n",index,value.name,value.id)
		}
	}
}
