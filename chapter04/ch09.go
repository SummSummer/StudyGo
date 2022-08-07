package main

func test() []func() {
	var s []func()
	// for循环复用局部变量i，每次添加匿名函数引用的都是同一个i
	for i := 0; i < 2; i++ {
		s = append(s, func() {
			println(&i, i)
		})
	}
	return s
}

func main() {
	// main函数执行这些函数时，读取的环境变量是i最后一次的值，因而循环输出的结果一样
	for _, f := range test() {
		f()
	}
}

// 如果想保证每个都不一致，可以使用不同的变量
func test01() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}

	return s
}

// 多个匿名函数引用同一个环境变量时，任何修改行为都会影响其它函数
func test02(x int) (func(), func()) {
	return func() {
		println(x)
		x += 10
	}, func() {
		println(x)
	}
}

func test03() {
	a, b := test02(100)
	a()		// 100
	b()		// 110
}