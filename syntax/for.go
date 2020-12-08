package main

import "fmt"

// go 语言中的循环结构只有 for 一种，for 循环的 {} 是必须的
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
