package main

import "fmt"

func intSeq() func() int {
	v := 0

	return func() int {
		v++
		return v
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
