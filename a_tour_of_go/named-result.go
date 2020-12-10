package main

import "fmt"

// 如果方法的返回结果与函数签名相符合，可省略返回时的参数
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
