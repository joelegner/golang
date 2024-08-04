package main

import (
	"math/rand"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Cycle through 8 billion people
const population = 10_000

const multiplier = 1.03

// const population = 8_000_000_000

// Cycle through 75 years' worth of days
const flipsPerLifetime = 100

const startingWealth = 1.0

type Person struct {
	id     int
	wealth float64
	heads  int
}

func main() {
	p := message.NewPrinter(language.English)
	var people [population]Person

	p.Printf("Cycling through %d people.\n", population)
	for personNumber := 0; personNumber < population; personNumber++ {
		people[personNumber] = Person{personNumber, startingWealth, 0}
		people[personNumber].wealth, people[personNumber].heads = simulateLifetime(people[personNumber], flipsPerLifetime)
		if personNumber%1000 == 1 {
			p.Printf("Person Number %d, wealth = %.10f, heads = %d\n", personNumber, people[personNumber].wealth, people[personNumber].heads)
		}
	}
	analyzeOutcomes(people)
}

func simulateLifetime(person Person, days int) (float64, int) {
	wealth := startingWealth
	heads := 0

	// Now we cycle through days. Each day the person gets
	// to flip a coin. If it is heads, they double their
	// current wealth. If it is tails, they halve it.
	for i := 0; i < days; i++ {
		if rand.Intn(2) == 1 {
			// You flipped heads! Congratulations!
			// You get to double your wealth.
			heads = heads + 1
			wealth = wealth * multiplier
		} else {
			// You flipped tails. Tough luck.
			// You have to halve your wealth.
			wealth = wealth / multiplier
		}
		// fmt.Println(i)

	}
	return wealth, heads
}

func analyzeOutcomes(people [population]Person) {
	var wealth [population]float64
	var maxWealth float64 = startingWealth
	var minWealth float64 = startingWealth
	var totalWealth float64 = startingWealth
	var maxHeads int = 0
	var minHeads int = 0
	var totalHeads int = 0

	for personNumber := 0; personNumber < population; personNumber++ {
		wealth[personNumber] = people[personNumber].wealth
		maxWealth = max(maxWealth, wealth[personNumber])
		minWealth = min(maxWealth, wealth[personNumber])
		totalWealth = totalWealth + wealth[personNumber]
		maxHeads = max(maxHeads, people[personNumber].heads)
		minHeads = min(maxHeads, people[personNumber].heads)
		totalHeads = totalHeads + people[personNumber].heads
	}

	averageWealth := totalWealth / population
	richestPercent := maxWealth / totalWealth * 100.0
	averageHeads := float64(totalHeads) / population

	p := message.NewPrinter(language.English)
	p.Printf("Total wealth = %.f\n", totalWealth)
	p.Printf("Maximum wealth = %.f (%.2f%%)\n", maxWealth, richestPercent)
	p.Printf("Average wealth = %.f\n", averageWealth)
	p.Printf("Minimum wealth = %.f\n", minWealth)
	p.Printf("Maximum heads = %.d\n", maxHeads)
	p.Printf("Average heads = %.3f\n", averageHeads)
	p.Printf("Minimum heads = %.d\n", minHeads)
	p.Printf("Gap = %.f\n", maxWealth/minWealth)
}
