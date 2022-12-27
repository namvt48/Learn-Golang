package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	// fmt.Println(num, err)
	// num = strings.Trim(num, "\r\n")
	num = strings.TrimSuffix(num, "\n")
	if number, err := strconv.ParseFloat(num, 64); err == nil {
		fmt.Println(number)
	}
}
