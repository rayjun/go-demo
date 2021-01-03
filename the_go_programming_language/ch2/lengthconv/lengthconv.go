package lengthconv

import "fmt"

type Meter float64
type Inch float64

func MToI(m Meter) Inch {
	return Inch(m * 39.37)
}

func IToM(i Inch) Meter {
	return Meter(i / 39.37)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (i Inch) String() string {
	return fmt.Sprintf("%gI", i)
}
