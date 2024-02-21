package shape

type Rectangle struct {
	sideA int
	sideB int
}

func (r Rectangle) Area() float64 {
	return float64(r.sideA * r.sideB)
}

func (r Rectangle) Type() string {
	return "Rectangle"
}

func CreateRectangle(sideA int, sideB int) Rectangle {
	return Rectangle{
		sideA: sideA,
		sideB: sideB,
	}
}
