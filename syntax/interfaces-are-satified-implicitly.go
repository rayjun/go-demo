package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 在这里 T 隐式的实现了接口 I，不需要用到 implements 关键字，这样就很好的将接口的定义和实现解耦了
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
