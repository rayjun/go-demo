package main

import "fmt"

// [n]T 表示一个包含 n 个 T 类型元素的数组，数组确定大小之后就不能改变
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
