package main

import (
	"fmt"
	"math"
)

var b int = 20

func main() {
	// var b, i, t = true, 20, "text"

	// fmt.Println(b, i, t)

	var x, y int = 3, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	fmt.Println(x, y, f)
}
