package main

import (
	"fmt"
	"strings"
)

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

	test_strings()

	test_pointers()

}

// Strings
func test_strings() {
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

	// String trimming using trim
	cutString := strings.Trim(ndString, "string")

	// String trimming using trimspace
	spaceString := "   trolled   "

	noSpaceString := strings.TrimSpace(spaceString)

	fmt.Println(noSpaceString)

	fmt.Println(cutString)

	// String splitting
	fullStr := "Hello there"

	splitStr := strings.Split(fullStr, " ")

	fmt.Println(splitStr)
}

func test_pointers() {
	// Pointers to variables
	testVal := 3

	pointVal := &testVal

	fmt.Println("Value of testVal: ", testVal)

	fmt.Println("Address of pointVal: ", pointVal)

	fmt.Println("Value of pointVal: ", *pointVal)

	// Pointers to functions
	var x = 200

	pointer_function(&x)

}

func pointer_function(a *int) {
	*a = 100
}

func test() {
	fmt.Println("Trolled")

}
