package main

import (
	"bufio"
	"fmt"
	"os"
)

// Output lines that have duplicates through stdin
func main() {
	// Create new empty map of string key and int value
	counts := make(map[string]int)

	// Create a Scanner variable
	// Scanners reads input and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)

	// While loop to go through all input from stdin
	// Scan() method returns true if there is line, false otherwise
	for input.Scan() {
		// The scanned line, input.Text() is used as a key to counts
		// Increment the value of the key by one, because default is zero
		counts[input.Text()]++
	}

	for key, val := range counts {
		if val > 1 {
			// C-style printf
			// In Go, these conversions (%d, %s, ...) are called verbs
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}
