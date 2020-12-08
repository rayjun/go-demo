package main

import "fmt"

// go 语言的函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
