package main

import "fmt"

func showType(v any) {
	fmt.Printf("%v - %T\n", v, v)
}

func arrays() {
	var students [3]string
	students[0] = "Lucas"
	showType(students)

	grades := [...]int{97, 85, 93}
	showType(grades)
	showType(len(grades))

	matrix := [...][3]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	showType(matrix)

	copy := matrix
	copy[0][0] = 0
	showType(copy)

	reference := &matrix
	(*reference)[1][1] = 1
	showType(matrix)
	showType(*reference)
}

func slices() {
	// Declaring and properties
	a := []int{1, 2, 3}
	showType(a)
	showType(len(a))
	showType(cap(a))

	// Referencing
	reference := a // Copies slice header, but points to same array
	showType(reference)
	showType(&a == &reference)       // Different, since they are different headers
	showType(&a[0] == &reference[0]) // Equal, since they point to same array

	// Pushing new elements
	c := []int{}
	showType(c)
	c = append(c, 1) // Can be expensive, may reallocate memory if not sufficient capacity
	c = append(c, 2, 3, 4, 5, 6)
	c = append(c, []int{12, 32}...) // spread
	showType(c)

	// Removing elements
	b := []int{0, 1, 2, 3, 4, 5, 6}
	showType(b)
	i := 3
	b = append(b[:i], b[i+1:]...)
	showType(b)
}

func main() {
	slices()
}
