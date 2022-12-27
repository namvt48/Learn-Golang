package main

import "fmt"

func plus(num *int) {
	*num++
}

func main() {
	i := 42

	p := &i

	fmt.Println("Before : ", *p)

	plus(p)

	fmt.Println("After: ", *p)
}
