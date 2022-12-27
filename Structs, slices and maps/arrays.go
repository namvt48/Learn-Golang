package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [...]int{1, 4, 5}

	fmt.Println(reflect.TypeOf(a))

}
