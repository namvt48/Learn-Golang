package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do("Hello")
	do(123)
	do(1.2)
}
