package main

import "fmt"

/**
* Go 语言有指针，一个指针表示持有一个值的内存地址
* *T 表示指向一个 T 类型值的指针
* var p *int
* & 符号用来生成一个指向运算元的指针
* i := 42
* p = &i
* 使用 * 意味着直接操作指针指向的值
* fmt.Println(*p) // 通过指针 p 读取 i
* *p = 21 // 通过指针 p 设置 i 的值
**/
func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针来赋值
	fmt.Println(i)  // 查看 i 的新值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针来计算 j
	fmt.Println(j) // 查看 j 的新值
}
