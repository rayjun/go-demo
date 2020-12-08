package main

import "fmt"

const (
	// 创建一个将 1 右移 100 位的数，就是 1 后面跟了 100 个 0
	Big = 1 << 100
	// 将这个值左移 99 位
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
