package main

import "fmt"

func showType(v any) {
	fmt.Printf("%v - %T\n", v, v)
}

// Bool
func bools() {
	var b bool = true
	showType(b)
}

// Int
func integers() {
	// Declaration
	var i int = 42 // int8, int16, int32, int64
	showType(i)

	// Unsigned (only positive)
	var ui uint = 92 // uint8, uint16, uint32
	showType(ui)

	// Operations
	var (
		a int = 10 // 1010
		b int = 3  // 0011
	)

	// Math
	fmt.Println(a + b) // ADD -> 13
	fmt.Println(a - b) // SUBTRACT -> 7
	fmt.Println(a * b) // MULTIPLY -> 30
	fmt.Println(a / b) // INTEGER DIVISION -> 3
	fmt.Println(a % b) // MOD -> 1

	// Bit Operations
	fmt.Println(a & b)  // AND -> 0010
	fmt.Println(a | b)  // OR -> 1011
	fmt.Println(a ^ b)  // XOR -> 1001
	fmt.Println(a &^ b) // AND NOT -> 0100
	fmt.Println(b << 1) // BITSHIFT LEFT - Multiply by 2 -> 0110
	fmt.Println(b >> 1) // BITSHIFT RIGHT - Integer Division by 2 -> 0001
}

// Float
func floats() {
	// Declaration
	var i float32 = 42.23 // float64
	showType(i)

	// Exponential Notation
	var j float64 = 3.12e21
	showType(j)

	// Operations
	var (
		a float64 = 10.5 // 1010.10
		b float64 = 3.25 // 0011.01
	)

	// Math
	fmt.Println(a + b) // ADD
	fmt.Println(a - b) // SUBTRACT
	fmt.Println(a * b) // MULTIPLY
	fmt.Println(a / b) // DIVISION

	// No mod or bit operations for Float
}

// Complex
func complexes() {
	var n complex64 = 1 + 1i // complex128
	showType(n)

	// Operations
	var (
		a complex64 = 3 + 6i
		b complex64 = 2 + 1i
	)

	// Math
	fmt.Println(a + b) // ADD
	fmt.Println(a - b) // SUBTRACT
	fmt.Println(a * b) // MULTIPLY
	fmt.Println(a / b) // DIVISION

	// Take REAL or IMAGINARY part of a complex number into float
	showType(real(a))
	showType(imag(a))

	// Create complex from Float
	var c complex64 = complex(5.1, 2.9)
	showType(c)
}

// Strings
func strings() {
	// Declaring String
	var s string = "Some letters"
	showType(s)

	// Strings are arrays of uint8
	showType(s[0])
	showType(string(s[0]))

	// Showing strings by their bytes
	var b []uint8 = []byte(s)
	showType(b)
}

func main() {
	//bools()
	//integers()
	//floats()
	//complexes()
	strings()
}
