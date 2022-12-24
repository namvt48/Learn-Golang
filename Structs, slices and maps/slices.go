package main

import "fmt"

func main() {
	// a := [6]int{1, 2, 3, 4, 5, 6}

	// slice := a[1:]

	// fmt.Println(slice)

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:4]
	// fmt.Println(a, b)
	// fmt.Println(b[0])

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)
	// fmt.Printf("%T", b)

	// q := [6]int{2, 3, 5, 7, 11, 13}
	// // fmt.Println(q)
	// fmt.Printf("%T\n", q)

	// r := q[:]
	// // fmt.Println(r)
	// fmt.Printf("%T\n", r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// // fmt.Println(s)
	// fmt.Printf("%T\n", s)

	// a := [...]int{12, 78, 50}
	// b := []int{12, 78, 50}
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T", b)

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[0:0] // s = [], len(s) = 0, cap(s) = 6

	// fmt.Println(cap(s))

	s = s[0:4] // s = [2, 3, 5, 7], len(s) = 4, cap(s) = 6
	s = s[2:4] // s = [5, 7], len(s) = 2, cap(s) = 4, cap được tính từ vị trí số 2 trở đi
	s = s[0:4] // s = [5, 7, 11, 13], len(s) = 4, cap(s) = 4

	var slice []int
	fmt.Println(slice, len(slice), cap(slice))

	if slice == nil {
		fmt.Println("nil")
	}

}
