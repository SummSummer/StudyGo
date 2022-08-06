package main

// 使用defer定义延迟调用，无论函数是否出错，都能保证在结束前被调用
// 可以定义多个defer，按照FIFO顺序执行
func test(a, b int) {
	defer println("dispose...")

	println(a/b)
}

func main() {
	test(10, 0)
}