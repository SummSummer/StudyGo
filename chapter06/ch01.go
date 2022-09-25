package main

import (
	"fmt"
	"sync"
)

// 方法
// 可以为当前包，以及除接口和指针外的任何类型定义方法
type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}

// 方法不支持重载，如果方法内部不引用实例，可以省略参数名
func (N) test01() {
	println("hi");
}

// 可以像访问匿名字段成员那样调用其方法
type data struct {
	sync.Mutex
	buf [1024]byte
}

func test02() {
	d := data{}
	d.Lock()			// 编译器会处理为 sync.(*Mutex).Lock()
	defer d.Unlock()
}

// 方法也会有同名遮蔽问题，可以利用这一特性，实现类似于override操作
type user struct{}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + "; manager"
}

func test03() {
	var m manager
	println(m.toString())
	println(m.user.toString())
}

func main() {
	
}