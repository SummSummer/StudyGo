package main

import "log"

// panic recover
// panic 会立即中断当前的函数流程，执行延迟调用，而在延迟调用中，recover可捕获并返回panic提交的错误对象

func main() {
	
}

func test01() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("i am dead")

	println("exit.")
}