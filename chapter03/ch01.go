package main

import "fmt"

// 除位移操作外，操作数类型必须相同，如果其中一个是无显式类型声明的常量，那么该常量操作数会自动转型
func test01() {
	const v = 20

	var a byte

	b := a + v							// v自动转为byte
	fmt.Printf("%T, %v\n", b, b)
	const c float32 = 1.2
	d := c + v							// v自动转为float32
	fmt.Printf("%T, %v\n", d, d)
}

// 右移操作数必须是无符号整数，或可以转换的无显式类型常量
func test02() {
	b := 23
	x := 1 << b
	println(x)
}

func main() {
	test01()
	test02()
}