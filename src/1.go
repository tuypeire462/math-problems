package main

import "math"

func solveMathProblem() {
	// Generate a random math problem
	a := rand.Intn(10) // Random number between 0 and 9
	b := rand.Intn(10) // Random number between 0 and 9
	op := ""           // Operator (+, -, *, /)
	switch {
	case rand.Intn(2) == 0:
		op = "+"
	default:
		op = "-"
	}

	// Print the problem and ask for input
	fmt.Printf("What is %d %s %d?\n", a, op, b)
	fmt.Scan(&answer)

	// Check if the answer is correct
	if answer == a+b {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect, try again.")
	}
}
