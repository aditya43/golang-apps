package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello Aditya, Namaste Aditya"

	// Replace 'Aditya' one time
	newStr := strings.Replace(str, "Aditya", "Nishi", 1)

	fmt.Println(newStr)
}
