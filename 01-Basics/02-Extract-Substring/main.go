package main

import "fmt"

func main() {
	str := "Namaste Aditya"
	greet := str[:7]
	name := str[8:]

	fmt.Println(greet, name)
}
