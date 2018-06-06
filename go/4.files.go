package main

import (
	"bufio"
	"fmt"
	"os"
)

// Output lines that have duplicates through stdin
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, filename := range files {
		// Returns open file (*os.File) used in subsequent reads
		// Second return value is a built-in error type
		f, err := os.Open(filename)

		// If error is not nil (None in Python), something is wrong
		if err != nil {
			// Print to stderr in its default format
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		// Call function, function declaration may be below, unlike C++
		countLines(f, counts)
		f.Close()
	}

	for key, val := range counts {
		if val > 1 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}

// Map all lines in file to its frequency
// The map made by make is passed by reference
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
