package shape

import "math"

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Type() string {
	return "Circle"
}

func CreateCircle(radius float64) Circle {
	return Circle{radius: radius}
}
