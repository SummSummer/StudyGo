package main

import "fmt"
// 让编译器检查，确保类型实现了指定接口
type x int

func init() {
	var _ fmt.Stringer = x(0)
}

// 定义函数类型，让相同签名的函数自动实现某个接口
type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string {
		return "Hello World!"
	})

	fmt.Println(t)
}