package main

import (
	"errors"
	"log"
)

// 异常处理

func main() {
	
}

// 官方推荐返回error状态，error作为最后一个参数

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}

	return x / y, nil;
}

func test01() {
	z, err := div(5, 0)
	if err == errDivByZero {
		log.Fatalln(err)
	}

	println(z)
}