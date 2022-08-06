package main

import (
	"errors"
	"fmt"
)

// 函数可以定义多个返回值，甚至可以对其命名
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a/b, nil
}

func main() {
	a, b := 10, 2
	c, err := div(a, b)
	
	fmt.Println(c, err)
}