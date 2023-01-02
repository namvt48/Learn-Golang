package main

import "fmt"

type I interface {
	M()
}

type S struct {
	name string
}

func (s S) M() {
	fmt.Println(s.name)
}

func main() {
	var i I = S{"Hello world"}
	i.M()
}
