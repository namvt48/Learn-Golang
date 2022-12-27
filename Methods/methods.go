package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	name string
	X, Y float64
}

type Student struct {
	name string
}

type myFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) getX(f float64) float64 {
	return v.X + f
}

func main() {
	v := Vertex{"A", 3, 4}
	fmt.Println(v.getX(1.0))
}
