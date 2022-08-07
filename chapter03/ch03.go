package main

import "fmt"

// for range，支持字符串，数组，数组指针，切片，字典，通道类型，返回索引、键值数据
func test01() {
	data := [3]string{"a", "b", "c"}

	for i, s := range data {
		println(i, s)
	}
}

// 可以使用 _忽略
func test02() {
	data := [3]string{"a", "b", "c"}

	for i := range data {
		println(i, data[i])
	}

	for _, i := range data {
		println(i)
	}

	for range data {}		// 仅迭代，不返回，可用来执行清空channel等操作
}

// range会复制目标数据，受直接影响的是数组，可改用数组指针或切片类型
func test03() {
	data := [3]int{10, 20, 30}

	for i, x := range data {		// 从data复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

	fmt.Println(data)

	for i, x := range data[:] {		// 仅复制slice，不包括底层array
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}

		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

	fmt.Println(data)
}

// 如果range目标表达式是函数调用，也仅被执行一次
func data() []int {
	println("origin data.")
	return []int{10, 20, 30}
}

func test04() {
	for i, x := range data() {
		println(i ,x)
	}
}

func main() {
	test04()
}