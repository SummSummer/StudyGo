package main

// 匿名函数

// 直接执行
func test01() {
	func(s string) {
		println(s)
	}("hello world!")
}

// 赋值给变量
func test02() {
	add := func(x, y int) int {
		return x + y
	}

	println(add(1, 2))
}

// 作为参数
func test03(f func()) {
	f()
}

func test04() {
	test03(func() {
		println("hello world!")
	})
}

// 作为返回值
func test05() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func test06() {
	add := test05()
	println(add(1, 2))
}

func main() {
	
}