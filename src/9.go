
func RandomMathProblem() {
	// Generate two random numbers between 1 and 10
	var num1 = rand.Intn(10) + 1
	var num2 = rand.Intn(10) + 1

	// Select a random operator (+, -, *, /)
	var op string
	switch rand.Intn(4) {
	case 0:
		op = "+"
	case 1:
		op = "-"
	case 2:
		op = "*"
	case 3:
		op = "/"
	}

	// Return the problem in the format of num1 op num2
	return fmt.Sprintf("%d %s %d", num1, op, num2)
}