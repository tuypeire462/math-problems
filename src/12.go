  
package main

import "fmt"

func randomInt(min int, max int) int {
    return min + (max-min)*rand.Float64()
}

func generateProblems() []problem {
	var problems = make([]problem, 0)
	for i := 0; i < 10; i++ {
		prob := problem{
			answer: randomInt(1, 20),
		}
		for j := 0; j < 4; j++ {
			prob.choices = append(prob.choices, randomInt(1, 20))
		}
		problems = append(problems, prob)
	}
	return problems
}

func main() {
	fmt.Println(generateProblems())
}