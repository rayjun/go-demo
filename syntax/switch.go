package main

import (
	"fmt"
	"runtime"
)

// 与其他语言一样，go 语言也有 switch，但 go 语言中的 switch 只会执行其中的一个分支，相当于每个分支后面自动执行了 break，而且 switch 的条件不要求是常量，其他的类型也是可以的
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
