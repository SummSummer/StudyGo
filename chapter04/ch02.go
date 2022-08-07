package main

// 定义函数类型，方便阅读和维护
type FormatFunc func(string, ...interface{}) (string, error)

func format(f FormatFunc, s string, a ...interface{}) (string, error) {
	return f(s, a...)
}

// 函数只能判断其是否为nil，不支持其他比较操作
func a() {}
func b() {}

func test01() {
	println(a == nil)
	// println(a == b)
}

// 从函数返回局部变量指针是安全的，编译器会通过逃逸分析来决定是否在堆上分配内存
func c() *int {
	a := 0x100
	return &a
}

func test02() {
	var a *int = c()
	println(a, *a)
}

func main() {
	
}