package shape

import "fmt"

type Shape interface {
	Area() float64
	Type() string
}

func PrintInfo(s Shape) {
	fmt.Printf("Shape information - name %v squre %v \n", s.Type(), s.Area())
}
