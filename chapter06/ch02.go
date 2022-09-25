package main

import "fmt"

// 表达式
type N int

func (n N) test01() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

// method expression
// 通过类型引用的method expression会被还原为普通函数，receiver是第一参数，调用时需显式传参
func test02() {
	var n N = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1 := N.test01
	f1(n)

	f2 := (*N).test01
	f2(&n)
}

// method value
// 通过实例或指针引用，会立即计算并复制该方法执行所需要的receiver对象，与其绑定，在之后的执行时，会隐式传入receiver参数
func test03() {
	var n N = 25

	p := &n

	n++
	f1 := n.test01

	n++
	f2 := p.test01

	n++
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1()
	f2()
}


func main() {
	
}