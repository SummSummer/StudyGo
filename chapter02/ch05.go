package main

import (
	"fmt"
	"math"
	"strconv"
)

// 支持八进制，十六进制和科学计数法
func main() {
	a, b, c := 100, 0144, 0x64

	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#x\n", a, a, a)

	fmt.Println(math.MinInt8, math.MaxInt8)
	test()
	test01()
}

// 标准库 strconv 可在不同进制(字符串)间转换
func test() {
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)

	println(a, b, c)

	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("0x" + strconv.FormatInt(a, 16))
}

// 浮点数需要注意小数位的有效精度
func test01() {
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	println(a, b, c)
	println(a == b, a == c)
	fmt.Printf("%v, %v, %v\n", a, b, c)
}