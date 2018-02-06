package main

import "fmt"
// Package is in current folder and imported with alias `asd`
import asd "./sub"

func main() {
	var input float64
	fmt.Scanln(&input)

	fmt.Println(asd.Doub(input))

	fmt.Scanln(&input)
}