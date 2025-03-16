  package main

import "math/rand"

func main() {
	// Generate a random number between 1 and 10
	x := rand.Intn(10) + 1

	// Generate a random number between 1 and 10
	y := rand.Intn(10) + 1

	// Calculate the sum of x and y
	sum := x + y

	// Print the result
	fmt.Println("The sum of", x, "and", y, "is", sum)
}