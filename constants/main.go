package main

import "fmt"

func showType(v any) {
	fmt.Printf("%v - %T\n", v, v)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	// Const declaration
	const i = 21
	showType(i)

	// Must be calculable at compile time
	// const b = add(2, 3)

	// Enum constants
	const (
		none = iota // Can also "_ = iota" to not store value of 0
		professor
		student
		principal
	)
	var role int = professor
	fmt.Printf("%v\n", role == professor)

	// Numerical patterns
	const (
		a = iota * 2
		b
		c
	)
	showType(a)
	showType(b)
	showType(c)

	// Store multiple flags
	const (
		canRead  = 1 << iota // 001
		canWrite             // 010
		canRun               // 100
	)
	var permissions byte = canRead | canRun
	fmt.Printf("%b\n", permissions)
	fmt.Printf("%v\n", permissions&canRead == canRead)
}
