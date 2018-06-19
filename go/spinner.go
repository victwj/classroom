package main

import (
	"fmt"
	"time"
)

// Iterates through spin frame every delay duration
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			// \r replaces prev character
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	// Start the spinner as a goroutine
	go spinner(100 * time.Millisecond)
	const n = 30
	time.Sleep(5 * time.Second)
}
