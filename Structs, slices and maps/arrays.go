package main

import "fmt"

func main() {
	var array [2]string
	array[0] = "hello"
	array[1] = "world"

	fmt.Println(array[0], array[1])
	fmt.Println(array)

	primes := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println(primes)

}
