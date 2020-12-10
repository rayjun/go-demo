package main

import "fmt"

// Go 中的函数可以用作闭包，闭包是指一个函数值引用着外部的变量，每个函数都可以访问和修改这个变量，这个变量和这个函数绑定在一起
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
