package main

import (
	"fmt"
	"math"
)

// if 和 for 类似，{} 是必须的，参数部分不需要 ()
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
