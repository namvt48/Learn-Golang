package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// fmt.Println(add(1, 2))
	// fmt.Println(swap(3, 7))
	fmt.Println(split(20))
}
