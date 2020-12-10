package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 相当于实现 Java 中的 toString 方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
