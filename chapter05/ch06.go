package main

import "fmt"

// 字典

func test01() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}

	fmt.Println(m, m2)
}

func test02() {
	m := map[string]int {
		"a": 1,
		"b": 2,
	}

	m["a"] = 10
	m["c"] = 30

	if v, ok := m["d"]; ok {
		println(v)
	}

	delete(m, "d")
}

// 对字典进行迭代时，每次返回的键值次序都不同
func test03() {
	m := make(map[string]int)

	for i := 0; i < 8; i++ {
		m[string('a'+i)] = i
	}

	for i := 0; i < 4; i++ {
		for k, v := range m {
			println(k, ":", v, " ")
		}

		println()
	}
}

// len返回当前键值对的数量，cap不接受字典类型，字典被设计成 not addressable ，不能直接修改value成员的值(数组或结构)
func test04() {
	type user struct {
		name string
		age int
	}

	m := map[int]user {
		1: {"Tom", 10},
	}

	// m[1].age += 1

	println(m)
}

// 可以通过返回整个value，修改后在设置字典的值，或直接使用指针
func test05() {
	type user struct {
		name string
		age int
	}

	m := map[int]user {
		1: {"Tom", 19},
	}

	u := m[1]
	u.age += 1
	m[1] = u

	m2 := map[int]*user {
		1: &user{"Jack", 10},
	}

	m2[1].age++
}

// nil字典和空字典
func test06() {
	var m map[string]int	// nil字典可读不可写
	println(m["a"])			// 读时返回零值
	m["a"] = 1				// 写时报错 panic: assignment to entry in nil map

	var m1 map[string]int	
	m2 := map[string]int{}	// m2已初始化，相当于make操作

	println(m1 == nil, m2 == nil)	// true false
}

func main() {
	test03()
}