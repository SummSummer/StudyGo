package main

func main() {

}

// 控制语句只有三种
func if_demo() {
	x := 100

	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println(0)
	}
}

func switch_demo() {
	x := 100

	switch {
		case x > 0:
			println("x")
		case x < 0:
			println("-x")
		default:
			println("0")
	}
}

func for_demo_01() {
	x := 0

	for x < 5 {
		println(x)
		x++
	}
}

func for_demo_02() {
	x := 0

	for {
		if x > 5 {
			break
		}
		println(x)
		x++
	}
}

// 使用 for range遍历时，可以同时返回索引和值
func for_demo_03() {
	x := []int{1, 2, 3}

	for i, n := range x {
		println(i, ":", n)
	}
}