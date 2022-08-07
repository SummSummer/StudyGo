package main

import "fmt"

// 使用type自定义类型
type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func test() {
	// 可以使用type语句块定义多个类型，在函数或代码块内定义局部类型
	type (
		user struct {
			name string
			age uint8
		}

		event func(string) bool
	)

	u := user{"Tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")
}


// 自定义类型和别名不同，无法进行隐式转换，不能直接用于比较表达式
func test01() {
	type data int
	var d data = 10

	var x int = int(d)			// 需要显式转换
	println(x)

	println(d == data(x))		// 需要显式转换
}

// struct tag属于类型组成部分，不仅仅是元数据描述
func test02() {
	var a struct {
		x int `x`
		s string `s`
	}

	var b struct {
		x int 
		s string 
	}

	// cannot use a (variable of type struct{x int "x"; s string "s"}) as struct{x int; s string} value in assignment
	// b = a

	fmt.Println(a, b)
}

// 函数的参数顺序也属于签名组成部分
func test03() {
	var a func(int, string)
	var b func(string, int)

	// cannot use a (variable of type func(int, string)) as func(string, int) value in assignment
	// b = a

	a(1, "s")
	b("s", 1)
}


func main() {
	f := read | exec
	fmt.Printf("%b\n", f)
}