package main

import "fmt"

//slice 功能上与数组类似，数组有固定的大小，而 slice 的大小是可以动态调整的，slice 使用的比数组更加广泛
// []T 就表示 T 类型的 slice
// clice 是一个半开范围，左闭右开，假如要获取一个数组中下标为 1~3 元素： a[1:4]
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
