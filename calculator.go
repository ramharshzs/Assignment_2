package main

import (
	"errors"
)

func calculator(operator string, num1, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("invalid operator")
	}
}
