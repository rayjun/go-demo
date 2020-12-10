package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T // 如果实现是 nil，那么下面的方法就会被 nil receiver 进行调用，在其他的语言中，这可能会触发空指针异常，但是在 Go 语言中不会
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}
