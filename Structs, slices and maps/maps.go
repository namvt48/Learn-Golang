package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["key1"] = 1
	fmt.Println(m["key1"])

	v0, ok0 := m["key1"]
	fmt.Println(v0, ok0)

	delete(m, "key1")
	fmt.Println(m["key1"])

	v, ok := m["key1"]
	fmt.Println(v, ok)

	v1, ok1 := m["key2"]
	fmt.Println(v1, ok1)

}
