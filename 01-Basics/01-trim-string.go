package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\t Namaste Aditya     "
	fmt.Printf("%d %s", len(str), str)

	trimmedStr := strings.TrimSpace(str)
	fmt.Printf("%d %s", len(trimmedStr), trimmedStr)
}
