package main

// 数组

// 初始化
func test01() {
	var a [4]int

	b := [4]int{2, 5}

	c := [4]int{2, 3: 5}

	d := [...]int{1, 2, 3}

	e := [...]int{10, 3: 100}

	println(a, b, c, d, e)
}

// Go数组是值类型，不会像C中数组隐式作为指针，赋值和传参操作都会复制整个数组数据

func main() {
	
}