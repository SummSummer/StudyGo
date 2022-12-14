package main

func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}

	x := calc {
		mul: func(x, y int) int {
			return x * y
		},
	}

	println(x.mul(2, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}

	println((<-c)(1, 2))
}

func test() {
	// 未被使用的匿名函数会报错
	// func(s string) {
	// 	println(s)
	// }
}

func main() {
	
}