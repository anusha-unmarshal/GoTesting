package length

const (
	In Length        = 2.54
	Centimeter Length = 1.
	Feet Length      = 30.48
	Meter      = 100* Centimeter
)

type Length float64

func (len Length) Meter() float64 {
	return float64(len/Meter)

}




