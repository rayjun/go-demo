package main

import (
	"bufio"
	"fmt"
	"go-demo/the_go_programming_language/ch2/lengthconv"
	"go-demo/the_go_programming_language/ch2/tempconv"
	"go-demo/the_go_programming_language/ch2/weightconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		conv(t)
	}

	if len(os.Args) <= 1 {
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			if input.Text() == "" {
				break
			}
			t, err := strconv.ParseFloat(input.Text(), 64)

			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			conv(t)
		}
	}
}

func conv(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	m := lengthconv.Meter(t)
	i := lengthconv.Inch(t)

	fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToI(m), i, lengthconv.IToM(i))

	k := weightconv.Kg(t)
	p := weightconv.Pound(t)

	fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
}
