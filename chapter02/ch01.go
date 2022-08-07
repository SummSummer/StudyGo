package main

var s string		// 全局变量可以只定义不使用，局部变量定义后必须使用

func main() {
	var x int		// 会自动初始化为零值
	var y = false	// 自动推断数据类型

	var (
		a, b = 100, "abc"
		c, d int
	)

	println(x, y, a, b, c, d)
}

func variable() {
	/*
		简短定义模式的限制
			定义变量，同时显示初始化
			不能提供数据类型
			只能在函数内部使用
	*/
	x := 100
	println(x)
}

func variable_01() {
	// 简短模式并不一定是定义变量，也可能是赋值操作。
	// 简短模式退化的前提：至少有一个新变量被定义，且必须在同一个作用域
	x := 100
	println(&x)

	x, y := 200, "abc"		// x为赋值操作

	println(&x)
	println(y)
}

func variable_02() {
	// 多变量赋值时，会先计算出所以右值，然后依次赋值
	x, y := 1, 2
	x, y = y+3, x+2
	println(x, y)
}