package main

import "lesson_5/shape"

func main() {
	c := shape.CreateCircle(10.0)
	r := shape.CreateRectangle(10, 10)

	shape.PrintInfo(c)
	shape.PrintInfo(r)
}
