package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (p User) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	user1 := User{"Nam", 12}

	fmt.Println(user1)
}
