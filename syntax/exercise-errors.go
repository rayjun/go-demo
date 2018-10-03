package main

import (
	"fmt"
	"math"
)


type ErrNefativeSqrt float64



func (e ErrNefativeSqrt) Error() string {
	return fmt.Sprintf("cant Sqrt negative number: %v", float64(e) )
}


func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNefativeSqrt(x)
	}

	
	return math.Sqrt(x), nil
}


func main () {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}