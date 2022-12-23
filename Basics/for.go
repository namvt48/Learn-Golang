package main

import "fmt"

func main() {
	sum := 1

	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	for sum < 1000 {
		// fmt.Println(sum)
		sum += sum
	}

	if sum%2 == 0 {
		fmt.Print("Chan")
	} else {
		fmt.Print("Le")
	}
	// fmt.Println(sum)
}
