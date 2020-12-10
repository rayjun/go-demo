package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // 会报运行时异常，接口必须要能绑定特定的类型才能进行方法调用
}

func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}
