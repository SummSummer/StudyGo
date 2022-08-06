package main

// 函数作为第一类型，可作为参数或返回值
func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	x := 100

	f := test(x)
	f()
}