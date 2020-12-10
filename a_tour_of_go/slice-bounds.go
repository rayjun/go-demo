package main

import "fmt"

// 在使用 slice 时，不用开始回想结束的位置不用全部指定
// 开始的位置默认为 0，结束的位置默认为 slice 的长度
// var a [10]int
// 对于上的数组，下面的slice 方式都是等价的：
// a[0:10]
// a[:10]
// a[0:]
// a[:]
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
