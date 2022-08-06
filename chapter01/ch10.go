package main

import "fmt"

// 可以直接调用匿名字段的方法，类似于继承
type user struct {
	name string
	age byte
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name = "Tom"
	m.age = 29

	println(m.ToString())		// 使用manager对象调用user的方法
}