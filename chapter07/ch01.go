package main

// 接口
// 只要目标类型方法集内包含接口的全部方法，就被视为实现了该接口
type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test() {}
func (data) string() string {
	return ""
}

func main() {
	var d data

	var t tester = &d

	t.test()
	println(t.string())
}

// 接口变量默认值为nil，如果实现接口的类型支持，可做相等运算
func test01() {
	var t1, t2 interface{}
	println(t1 == nil, t1 == t2)	// true true

	t1, t2 = 100, 100
	println(t1 == t2)	// true

	t1, t2 = map[string]int{}, map[string]int{}	
	println(t1 == t2)	// panic: error
}