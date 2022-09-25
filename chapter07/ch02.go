package main

// 接口可以嵌入其他接口，目标类型方法集中必须拥有包含嵌入接口方法在内的全部方法才算实现了该接口
type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct {}

func (*data) test() {}
func (data) string() string {return ""}

func main() {
	var d data
	var t tester = &d
	t.test()
	println(t.string())	
}

// 超集接口变量可隐式转换为子集，反之不行
func pp(a stringer) {
	println(a.string())
}

func test01() {
	var d data
	var t tester = &d
	pp(t)

	var s stringer = t
	println(s.string())

	// var s2 tester = s
}

// 支持匿名接口类型，可直接用于变量定义，或作为结构字段类型
type node struct {
	data interface {
		string() string
	}
}

func test02() {
	var t interface {
		string() string
	} = data {}

	n := node {
		data: t,
	}

	println(n.data.string())
}