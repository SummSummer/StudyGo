package main

// 结构struct

// 可以通过指针直接操作结构字段，但不能是多级指针
func test01() {
	type user struct {
		name string
		age int
	}

	p := &user {
		name: "Tom",
		age: 10,
	}

	p.name = "Jack"
	p.age++

	// p2 := &p
	// *p2.name = "Cat"
}

func main() {
	
}