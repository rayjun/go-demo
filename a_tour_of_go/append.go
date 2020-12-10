package main

import "fmt"

// slice 可以使用 append 来动态调整 slice，如果 slice 的 cap 不足以存放新的元素，就会创建一个新的数组，然后返回的 slice 指向新的数组
func main() {
	var s []int
	printSlice(s)

	// 向 nil 中追加元素
	s = append(s, 0)
	printSlice(s)

	// 追加单个元素
	s = append(s, 1)
	printSlice(s)

	// 可以一次追加多个元素
	s = append(s, 2, 3, 4, 5, 6, 7)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
