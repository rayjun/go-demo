package main

import "fmt"

// 一个函数可以有零个或者多个参数，函数的返回类型在方法签名的最后面
// 需要注意，变量的类型在变量后面

func add(x int, y int) int {
    return x + y
}


func main() {
    fmt.Println(add(42, 13))
}
