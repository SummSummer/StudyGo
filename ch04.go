package main

import "fmt"

// 可变参数本质是一个切片，只能接受一个或多个同类型的参数，且必须放在列表尾部
func test01(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a)
}

// 将切片作为变参时，须进行展开操作。如果是数组，先将其转换为切片
func test02(a ...int) {
	fmt.Println(a)
}

func test03() {
	a := [3]int{10, 20, 30}
	test02(a[:]...)
}

// 变参是切片，那么参数复制的仅是切片本身，并不包括底层数组，也因此可修改原数据。
// 如果需要，可用内置函数copy复制底层数据
func test04(a ...int) {
	for i := range a {
		a[i] += 100
	}
}

func test05() {
	a := []int{10, 20, 30}
	test04(a...)

	fmt.Println(a)
}

func main() {
	test01("abc", 1, 2, 3, 4)	
}