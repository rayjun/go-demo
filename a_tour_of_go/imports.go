package main

import (
    "fmt"
    "math"
)

//import "fmt"
//import "math"

// 使用 import 可以导入其他包的代码，可以每行导入一个代码，但比较推荐的方式使用 () 导入的所有包
// 不可以重复的导入包
// 已经导入的包必须要使用
func main() {
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(11))
}
