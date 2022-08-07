package main

// 闭包：通过指针引用环境变量，导致其生命周期延长

func test(x int) func() {
	println(&x)

	return func() {
		println(&x, x)
	}
}

func main() {
	f := test(123)
	f()	
}