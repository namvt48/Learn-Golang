package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var small float64 = 0.0001

	for math.Abs(x-z) > small {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

/*
	x = 2

z = 1 abs =
*/
func main() {

	// Sqrt(2)
	fmt.Println(math.Sqrt(2))

}
