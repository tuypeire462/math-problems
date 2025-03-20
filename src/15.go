package main
import (
"fmt"
"math/rand"
)
func main() {
const num1 = rand.Intn(9) + 1
const num2 = rand.Intn(9) + 1
if num1 > num2 {
fmt.Println("The solution is", num1 - num2)
} else if num1 == num2 {
fmt.Println("The solution is", 0)
} else {
fmt.Println("The solution is", num2 - num1)
}
}