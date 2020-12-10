package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Channels 的发送不再发送数据的时候，就可以关闭 channels，接收方可以通过 channels 的第二个参数来判断是否关闭
// v, ok := <-ch // 如果为 false，表示 channels 已经关闭
// channels 应该只由发送方来关闭
// 而且 channels 通常可以不用手动关闭
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
