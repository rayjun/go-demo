package main

import "fmt"

// 如果两个参数的类型相同，只需要在最后声明参数的类型就可以
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
