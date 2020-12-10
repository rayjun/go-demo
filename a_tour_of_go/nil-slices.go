package main

import "fmt"

// 不包含元素的 slice 是 nil，length 和capacity 都是 0
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
