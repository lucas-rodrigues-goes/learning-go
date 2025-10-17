package main

import (
	"fmt"
	"strconv"
)

// Variables
var Variable int = 10 // Declare globally
func variables() {
	var a int // Declare without initializing
	a = 10
	fmt.Println(a)

	var b int = 20 // Specify type
	fmt.Println(b)

	c := 30.2 // Automatically detect type
	fmt.Println(c)

	// Declare a block for multiple variables
	var (
		name string = "Lucas"
		age  int    = 22
	)
	fmt.Println(name, " ", age)

	var theURL string = "https://www.google.com" // Acronyms should be uppercase
	fmt.Println(theURL)
}

// Conversions
func conversions() {
	var i int = 42
	fmt.Printf("%v - %T\n", i, i)

	// Converting between number types
	var j float32 = float32(i)
	fmt.Printf("%v - %T\n", j, j)

	// Converting to strings
	var k string = strconv.Itoa(i)
	fmt.Printf("%v - %T\n", k, k)
}

func main() {
	//variables()
	//conversions()
}
