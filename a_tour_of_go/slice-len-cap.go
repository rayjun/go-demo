package main

import "fmt"

// slice 有两个属性，length 和 capacity
// slice 的 length 表示其中包含的元素个数，使用 len(s)
// slice 的 capacity 表示内部数组的元素个数，使用 cap(s) 获得
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice 的长度为0
	s = s[:0]
	printSlice(s)

	// 扩充长度
	s = s[:4]
	printSlice(s)

	// 删除开始的两个元素
	s = s[2:]
	printSlice(s)

	s = s[:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
