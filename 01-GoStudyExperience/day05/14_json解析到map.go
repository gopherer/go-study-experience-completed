 package main 

 import (
	"encoding/json"
	"fmt"
)

 func main14(){
	jsonbuf := `
	{
		"company":"itcast",
		"subjects":[
			"Go",
			"C++",
			"Python",
			"Test"
		],
		"isok":true,
		"price":666.6666
	}`
	//创建一个map
	m := make(map[string]interface{},4)

	err := json.Unmarshal([]byte(jsonbuf),&m)//第二个参数要传递地址
	if err != nil{
		fmt.Println("err = ",err)
		return 
	}
	fmt.Printf("m = %+v\n",m)
	
// 	var str string
// 	str = string(m["company"]) err无法转换
// 	fmt.Println("str = ",str)
	
	//类型断言 来实现转换
	for key, value := range	m{
		//fmt.Printf("%v =============%v\n",key,value)
		switch data := value.(type){
		case string:
			fmt.Printf("map[%s]的值类型为string, value = %s\n",key,data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool value = %v\n",key,value)
		case float64:
			fmt.Printf("map[%s]的值类型为float64, value = %f\n",key,data)
		//case []string: 无法匹配
		case []interface{}:
			fmt.Printf("map[%s]的值类型为string, value = %v\n",key,data)
		}
	}
}