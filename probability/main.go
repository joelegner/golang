package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This mysterious code is how you seed the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Set up a counter for hits
	hits := 0

	// 1:x odds
	var one_in int64 = 365 * 30

	fmt.Println("Hello, Go!")
	for k := 0; k <= 30; k++ { // Cycle through 30 years
		for i := 0; i < 365; i++ { // Cycle 365 days
			if r.Int63n(one_in) == 1 { // Check for a hit
				hits++ // Count a hit
			}
		}
	}
	fmt.Printf("%d hits\n", hits)
}
