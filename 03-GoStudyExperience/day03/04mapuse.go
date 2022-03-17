package main

import "fmt"

func main() {
	StdMap := make(map[string]map[string]string)

	StdMap["stu01"] = make(map[string]string)
	StdMap["stu01"]["name"] = "tom"
	StdMap["stu01"]["sex"] = "男"
	StdMap["stu01"]["addr"] = "福建"

	StdMap["stu02"] = make(map[string]string)
	StdMap["stu02"]["name"] = "Jacki"
	StdMap["stu02"]["sex"] = "女"

	fmt.Println(StdMap)
	fmt.Println(StdMap["stu01"])
	fmt.Println(StdMap["stu02"])
}
