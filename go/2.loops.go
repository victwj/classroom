package main

import (
	"fmt"
	"os"
)

// Echoes command line arguments 3x with different loops
func main() {
	// Initialize two variables s and sep
	// If not initialized, implicitly initialized to zero value for its type

	var s, sep string
	sep = " "

	/* Different variable declarations
	s := ""
		- Most compact, but can only be used within a function, not package-level
	var s string
		- Uses default initialization to the zero value
	var s = ""
		- Used when declaring multiple variables
	var s string = ""
		- Like above, but for potentially multiple variables with different types
	*/

	// For loop, := is a short variable declaration
	// It declares variables and gives appropriate types based on init values
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
	}

	// While loop, no initialization and no post
	i := 1
	for i < len(os.Args) {
		s += os.Args[i] + sep
		i++
	}

	// Range based for loop
	// The range command produces a pair of values, the index and value of element at index
	// We don't need index, so put it in a blank identifier _
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}

	// Infinite loop
	// for {}

	fmt.Println(s)
}
