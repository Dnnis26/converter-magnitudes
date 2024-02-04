package converter

type Unit float64

func MtoF(meters float64) Unit {
	return Unit(meters / 0.3048)
}

func FtoM(feet float64) Unit {
	return Unit(feet * 0.3048)
}
