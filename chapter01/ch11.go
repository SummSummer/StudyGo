package main

import "fmt"

// 接口采用duck type方式，无须在实现类型上添加显示声明
type user struct {
	name string
	age byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {		// 定义接口
	Print()
}

func main() {
	var u user
	u.name = "Tom"
	u.age = 29

	var p Printer = u			// 只要包含接口所需要的全部方法，就表示实现了该接口
	p.Print()
}