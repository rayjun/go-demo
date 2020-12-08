package main

import "fmt"

// := 直接赋值只能在方法体内使用，在方法体外无法使用

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
