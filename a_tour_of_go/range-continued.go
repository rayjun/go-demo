package main

import "fmt"

// 可以通过 _ 来省略下标或者值
// 如果只想要下标，可以直接使用 :
// i := range pow
func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = i << uint(i) // 2**i
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
