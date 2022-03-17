package main

import (
	"fmt"
	"regexp"
)

func main10() {
	buf := `

<div>哈哈哈</div>

<div>哈哈哈
	你瞅啥
</div>
}

	`
	reg := regexp.MustCompile(`<div>(?s:(.*?))<\/div>`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	result := reg.FindAllStringSubmatch(buf, -1)
	//fmt.Println("result = ", result)

	//过滤<> </>
	for _, test := range result {
		//fmt.Println("test[0] = ", test[0])//带<></>
		fmt.Println("test[1] = ", test[1]) //不带<></>
	}

}
