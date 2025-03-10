package main

import "fmt"

func main() {
    // Generate a random number between 1 and 10
    randNum := 1 + rand.Intn(9)
    fmt.Println("The random number is", randNum)
}
