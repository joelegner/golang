package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomInt := rand.Intn(100) + 1
	fmt.Println(randomInt)
}
