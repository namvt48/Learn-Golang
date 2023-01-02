package main

import "fmt"

func plus(num *int) {
	*num++
}

type O func(a, b int) int

type B int

func main() {
	// i := 42

	// p := &i

	// b := new(int)
	// *b = 41

	// a := 10
	// fmt.Println("Before : ", a)

	// plus(&a)

	// fmt.Println("After: ", a)
	var b B
	b = 10
	fmt.Println(b)
}
