package main

import "fmt"

// for 循环可以使用如下的写法，这样就类似于其他语句中的 while 了
func main() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
