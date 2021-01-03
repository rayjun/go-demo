package weightconv

import "fmt"

type Kg float64
type Pound float64

func KToP(k Kg) Pound {
	return Pound(k * 2.205)
}

func PToK(p Pound) Kg {
	return Kg(p / 2.205)
}

func (k Kg) String() string {
	return fmt.Sprintf("%gKg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%gP", p)
}
