package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Print("Go runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	fmt.Printf("%s.\n", os)

	// }

	// today := time.Now().Weekday()
	// fmt.Println(today)
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("today")
	// case today + 1:
	// 	fmt.Println("1 day left")
	// case today + 2:
	// 	fmt.Println("2 days left")
	// default:
	// 	fmt.Println("default")
	// }

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
