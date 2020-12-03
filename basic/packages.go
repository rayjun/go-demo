package main

import (
    "fmt"
    "math/rand"
)

// 每个Go 程序都是由多个包来组成，程序会从名称为 main 的包中启动，如果启动的程序包名不是 main，程序无法运行
// 按照惯例，包的名称是导入路径的最后一个元素，比如 math/rand 包中的文件有着相同的包名 rand

func main() {
    fmt.Println("My favorite number is ", rand.Intn(10))
}
