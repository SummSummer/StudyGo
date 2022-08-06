package main

// 结构体struct，可以匿名嵌入其他类型
type user struct {
	name string
	age byte
}

type manager struct {
	user				// 嵌入其他匿名构造体
	title string
}

func main() {
	var m manager

	m.name = "Tom"		// 直接访问匿名字段的成员
	m.age = 29
	m.title = "CTO"

	println(m)
}