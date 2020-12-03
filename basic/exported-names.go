package main

import (
    "fmt"
    "math"
)

// 从包中导出的内容都是以大写字母开头，小写字母的表示不是导出的内容
func main() {
    fmt.Println(math.Pi)
}
