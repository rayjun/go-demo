package main

import (
	"fmt"
)


type Person struct {
	Name string
	Age int
}

/*
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
*/

func main(){
	a := Person{"AXX", 42}
	z := Person{"Zbbox", 9001}

	fmt.Println(a,z)
}