package main

import "fmt"

type Number float64
type User struct {
	Name string
	Age  int
}
type Increase interface {
	Plus() Number
}

func (num Number) Plus() Number {
	num++
	fmt.Print(num)

	return num
}

func (u User) String() string {
	return fmt.Sprintf("New user with name: %v and age: %v", u.Name, u.Age)
}

func main() {

	// user := User{
	// 	Name: "Nam",
	// 	Age:  20,
	// }

	// fmt.Println(user)

	var i Increase

	i = Number(1.2)
	i.Plus()
	fmt.Println(i)

}
