package main

import "errors"

// 命名返回值
// func paging(sql string, index int) (count int, pages int, err error) {}

// 命名返回值和参数一样，可当做函数局部变量使用，最后由return隐式返回
func div(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
		return
	}

	z = x/y
	return
}

// 如果代码块中定义了同名的变量，需要显式返回
func add(x, y int) (z int) {
	{
		z := x + y
		return z
	}
}

// 需要对所有的返回值命名：编译器会跳过未命名的返回值
// func test() (int, s string, e error) {
// 	return 0, "", nil
// }

// 如果返回值类型能够明确含义，尽量不使用返回值命名
// func NewUser() (*User, error) 

func main() {
	
}