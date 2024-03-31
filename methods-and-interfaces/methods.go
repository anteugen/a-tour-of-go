package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Circle struct {
	Radius float64
}

func setRadiusWithRegularFunc(c Circle, newRadius float64) Circle {
	c.Radius = newRadius
	return c
}

func setRadiusWithPointerFunc(c *Circle, newRadius float64) {
	c.Radius = newRadius
}

func (c Circle) setRadiusWithValueReceiver(newRadius float64) {
	c.Radius = newRadius
}

func (c *Circle) setRadiusWithPointerReceiver(newRadius float64) {
	c.Radius = newRadius
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{2, 4}
	fmt.Println(v)
	fmt.Println(v.Abs())

	c := Circle{Radius: 5.0}
	fmt.Println("Original radius:", c.Radius)

	c = setRadiusWithRegularFunc(c, 10.0)
	fmt.Println("Radius after being changed with a regular function:", c.Radius)

	setRadiusWithPointerFunc(&c, 15.0)
	fmt.Println("Radius after being changed with a regular function using a pointer:", c.Radius)

	c.setRadiusWithValueReceiver(20.0)
	fmt.Println("Radius after being changed with a value receiver:", c.Radius)

	c.setRadiusWithPointerReceiver(25.0)
	fmt.Println("Radius after being changed with a pointer receiver:", c.Radius)
}

