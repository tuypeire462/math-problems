// Package main implements a program to solve math problems.
package main

import "math"

// Problem represents a math problem that can be solved.
type Problem struct {
	Description string
	Answer      int
}

// Solve returns the solution to the problem.
func (p *Problem) Solve() int {
	return p.Answer
}

// Generate creates a new math problem and its solution.
func Generate(difficulty int) (*Problem, error) {
	switch difficulty {
	case 1: // Easy problems
		return &Problem{
			Description: "5 + 3",
			Answer:      8,
		}, nil
	case 2: // Medium problems
		return &Problem{
			Description: "10 - 4",
			Answer:      6,
		}, nil
	default: // Hard problems
		return &Problem{
			Description: "15 * 3",
			Answer:      45,
		}, nil
	}
}
