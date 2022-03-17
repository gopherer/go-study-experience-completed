package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//	Number1 即 multiplicand 被乘数  B
//	Number2 即 multiplier    乘数   C

func (user *User) Multiplication() string {
	if err := Clear(); err != nil {
	}
	mndF, merF := 0, 0
	//去掉点号，转为切片
	mndTemS := strings.Split(user.Number1, ".")
	merTemS := strings.Split(user.Number2, ".")
	//切片转字符串
	mndStr := mndTemS[0] + mndTemS[1]
	merStr := merTemS[0] + merTemS[1]
	//计算字符串的长度
	mndLen := len(mndStr)
	merLen := len(merStr)
	//
	ASlice := make([]int, 0)
	mndSlice := make([]int, mndLen)
	merMap := make(map[int]int)
	//将mndSte字符串转变为切片存入mndSlice
	for i := 0; i < mndLen; i++ {
		if mndStr[i] == 49 { //49 -- '1'
			mndSlice[i] = 1
		} else if mndStr[i] == 48 { //40 -- '0'
			mndSlice[i] = 0
		} else {
			if err := Clear(); err != nil {
			}
			GotPosition(X, Y)
			fmt.Print("输入有误,请重新输入")
			os.Exit(1)
		}
	}
	//处理mndSlice中的符号问题
	if mndSlice[0] == 1 {
		mndSlice[0] = 0
		mndF = 1
	}
	//将merStr字符串转变为Map存入merMap
	for i := 0; i < merLen; i++ { //去掉multiplier 中的都好后存入merMap 中
		if merStr[i] == 49 {
			merMap[i] = 1
		} else if merStr[i] == 48 {
			merMap[i] = 0
		} else {
			if err := Clear(); err != nil {
			}
			GotPosition(X, Y)
			fmt.Print("输入有误,请重新输入")
			os.Exit(1)
		}
	}
	//记录merMap中符号问题
	if merMap[0] == 1 {
		merF = 1
	}
	reverseSlice(mndSlice)
	for i := merLen - 1; i > 0; i-- {
		flag := 0
		if merMap[i] == 1 { // A<- 1/2(A+B) C-<1/2C
			if i == merLen-1 {
				ASlice = append(mndSlice, 0) //1/2B
			} else {
				for j := 0; j < len(ASlice); j++ { //A+B
					for {
						if len(ASlice) == len(mndSlice) {
							//ASlice[j] = (mndSlice[j] + ASlice[j] + flag) % 2
							//flag = (mndSlice[j] + ASlice[j] + flag) / 2
							tem := mndSlice[j] + ASlice[j] + flag
							if tem == 0 {
								ASlice[j] = 0
								flag = 0
							} else if tem == 1 {
								ASlice[j] = 1
								flag = 0
							} else {
								ASlice[j] = 0
								flag = 1
							}
							break
						} else {
							reverseSlice(mndSlice)
							mndSlice = append(mndSlice, 0)
							reverseSlice(mndSlice)
						}
					}
				}
				ASlice = append(ASlice, 0) //1/2(A+B)
			}
		} else { // i == 0
			ASlice = append(ASlice, 0)
		}
	}
	reverseSlice(ASlice)
	//结构符号处理
	if mndF == 1 || merF == 1 {
		if mndF+merF != 2 {
			ASlice[0] = 1
		}
	}
	//将int切片转为字符串
	resultStr := ""
	for i := 0; i < len(ASlice); i++ {
		if i == 1 {
			resultStr += "."
		}
		resultStr += strconv.Itoa(ASlice[i])
	}
	return resultStr
}
