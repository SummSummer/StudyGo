package main

import "fmt"

// 切片slice实现动态数组功能
func main() {
	x := make([]int, 0, 5)			// 创建容量为5的切片

	for i := 0; i < 8; i++ {
		x = append(x, i)			// 当容量不足时会自动扩容
	}

	fmt.Println(x)
}