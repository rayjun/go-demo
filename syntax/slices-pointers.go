package main

import "fmt"

// slice 就像是对数组的引用，实际上，slice 中不存储任何数据，只是表示数组的一部分，
// 如果对 slice 中的元素进行修改，那么数组的其他部分也会有变化
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}
