package challenges

func findFactorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * findFactorial(number-1)
}
