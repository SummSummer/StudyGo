package main

import "errors"

// _ 可作为忽略占位符，可作为表达式的左值，但无法读取内容
func main() {
	x, _ := test()		// 忽略函数返回的error信息
	println(x)
}

func test() (int, error){
	return 10, errors.New("hello")
}