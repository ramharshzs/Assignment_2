package main

func Sum(n int) int {
	if n <= 0 {
		return 0 // Return 0 if n is non-positive
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
	// Using the arithmetic sum formula
}
