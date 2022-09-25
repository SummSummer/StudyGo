package main

import (
	"log"
	"os"
)

// 延迟调用
func main() {
	
}

func test01() {
	f, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
}

// 延迟调用注册的是调用，必须提供执行所需的参数(哪怕为空)。参数值在注册时被复制并缓存。如果对状态敏感，可使用指针或闭包
func test02() {
	x, y := 1, 2

	defer func(a int) {
		println("defer x, y = ", a, y)		// y是闭包引用
	}(x)

	x += 100
	y += 200

	println(x, y)
}

// 多个延迟注册按FILO次序执行
func test03() {
	defer println("a")
	defer println("b")
}

func test04() (z int) {
	defer func() {
		println("defer: ", z)
		z += 100
	}()

	return 100			// 执行顺序：z = 100 -> call defer -> return
}

func test05() {
	println("test04:", test04())
}