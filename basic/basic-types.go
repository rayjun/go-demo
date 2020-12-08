package main

import (
	"fmt"
	"math/cmplx"
)

// go 中 的基本类型有：
// bool

// string

// 对于整数类型，如果没有特别的要求，比如限定大小或者要使用无符号的整数，否则直接使用 int 就可以
// int int8 int16 int32 int64
// uint uint8 uint 16 uint32 uint64 uintptr

// byte // uint8 的别名
// rune // int32 的别名，代表Unicode 编码

// float32 float64

// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
