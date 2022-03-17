package unit_testing

import "testing"

//注意文件名要以_test.go结尾
//函数名以TestXxx格式
func TestAdd(t *testing.T) {
	res := addTest(10, 20)
	if res != 30 {
		t.Fatalf("预期值 = %v 实际值 = %v", 30, res)
	}
}
func TestSub(t *testing.T) {
	res := subTest(30, 20)
	if res != 10 {
		t.Fatalf("预期值 = %v 实际值 = %v", 10, res)
	}
}
