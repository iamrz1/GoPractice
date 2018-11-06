package main

import "fmt"
import "math"

type shapes struct {
	color string
}
type circles struct {
	shapes
	radius float64
}

func (c *circles) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	fmt.Println("Program Started")

	shape := shapes{"blue"}
	//fmt.Println(shape.color)
	// circle := new(circles)

	// circle.color = "green"

	circle := circles{shape, 10}

	fmt.Println(circle.color)
	fmt.Println(circle.area())

	hain()
}
