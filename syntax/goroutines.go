package main

import (
	"fmt"
	"time"
)

// goroutine 是 Go 运行时的轻量级线程管理工具
// go f(x, y, z) 会开启一个新的 goroutine 运行
// goroutines 运行在相同的地址空间，如果访问内存必须是同步的
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
