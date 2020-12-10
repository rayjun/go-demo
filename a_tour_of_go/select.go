package main

import "fmt"

// select 可以让 goroutine i处理多个 channels
// select 块中只会有一个 case 被处理，会从中随机选择一个已经就绪的 case 进行处理
func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
