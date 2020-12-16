package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64 // exercise2.1

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Kelvin  = 0
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

// exercise2.1
func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
