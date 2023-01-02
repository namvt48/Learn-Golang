package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	s, oke := i.(string)
	fmt.Println(s, oke)

	f := i
	fmt.Println(f)

	fl, _ := i.(float64)
	fmt.Println(fl)
}
