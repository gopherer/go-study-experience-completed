package unit_testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (monster *Monster) write() bool {
	var filePath = "C:/GoCode/src/03_gocode/day05_file/test.txt"
	b, err := json.Marshal(monster)
	if err != nil {
		fmt.Print("json err = ", err)
		return false
	}
	fmt.Println(string(b))
	err = ioutil.WriteFile(filePath, b, 0666)
	if err != nil {
		fmt.Print("write file err ", err)
		return false
	}
	return true
}
func (monster *Monster) read() bool {
	var filePath = "C:/GoCode/src/03_gocode/day05_file/test.txt"
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print("read file err ", err)
		return false
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b, monster)
	if err != nil {
		fmt.Print("json err = ", err)
		return false
	}
	return true
}
func TestMonster(t *testing.T) {
	var monster = Monster{
		Name:  "猪八戒",
		Age:   600,
		Skill: "九齿钉耙",
	}
	sign := monster.write()
	if !sign {
		t.Fatalf("预期结果为%v 实际结果为%v", true, sign)
	}
	sign = monster.read()
	if !sign {
		t.Fatalf("预期结果为%v 实际结果为%v", true, sign)
	}
}
