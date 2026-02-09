package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Wait for 'GO' and press Enter as fast as you can...")
	rand.Seed(time.Now().UnixNano())
	delay := time.Duration(rand.Intn(3000)+2000) * time.Millisecond // Random delay between 2-5 seconds
	time.Sleep(delay)

	fmt.Println("GO!")
	start := time.Now()

	var input string
	fmt.Scanln(&input)

	elapsed := time.Since(start)
	fmt.Printf("Reaction time: %v milliseconds\n", elapsed.Milliseconds())
}
