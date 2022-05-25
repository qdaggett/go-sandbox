package main

import "fmt"

func main() {

	var X uint8 = 225

	fmt.Println(X, X-3)

	var Y int16 = 32767

	fmt.Println(Y+2, Y-2)

	// **Bitwise operations**
	p := 34
	q := 20

	// Bitwise and
	result1 := p & q
	fmt.Println(result1)

	// Misc operators
	a := 4

	b := &a // Address of a
	fmt.Println(*b)
	*b = 7 // Pointer to b
	fmt.Println(a)

	test()

	strings()

}

// Strings
func strings() {
	shString := "Shorthand string" // Shorthand declaration

	var ndString string
	ndString = "Normally declared string" // Regular string declaration

	for index, s := range "Test" { // Iterating over strng
		fmt.Println(index)
		fmt.Println(s)
	}

	// Printing
	fmt.Println(shString)
	fmt.Println(ndString)

}

func test() {
	fmt.Println("Trolled")

}
