package main

import "fmt"

// 类型转换

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

func test01() {
	var d data = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok {
		fmt.Println(n)
	}

	if d2, ok := x.(data); ok {
		fmt.Println(d2)
	}

	e := x.(error)
	fmt.Println(e)
}

func main() {
	var x interface{} = func(x int) string {
		return fmt.Sprintf("d:%d", x)
	}

	switch v := x.(type) {
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		println(v(100))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		println("unknow")
	}
}