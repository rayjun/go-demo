package main

import (
	"fmt"
	"time"
)

// 如果 switch 不用条件，则等同于 switch true，使用这个可以简化大段的 if-then-else 代码
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good eventing")
	}
}
