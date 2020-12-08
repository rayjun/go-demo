package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 当使用 range 来遍历 slice 时，第一个返回的是元素的下标，第二个是值
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d= %d\n", i, v)
	}
}
