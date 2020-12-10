package main

import "fmt"

// Channels 是可以被缓冲的，通过 make 的第二个参数就可以提供缓冲区的长度
// 当缓冲区满的时候，会被阻塞，缓冲区空的时候，就可以继续接收新的数据
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 这行代码写在这里会造成死锁，为什么？
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
