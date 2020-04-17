package main

import (
	"fmt"
)

func main() {
	var (
		oper string
		num1 float64
		num2 float64
	)
	fmt.Println("Enter the First Number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second Number:")
	fmt.Scanln(&num2)
	fmt.Println("Enter the Operation to perform:")
	fmt.Scanln(&oper)

	fmt.Println("First number is:", num1)
	fmt.Println("Second number is:", num2)
	result := process(oper, num1, num2)
	fmt.Println("Total:", result)
}

func process(oper string, num1, num2 float64) float64 {
	var (
		result float64
	)

	switch oper {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Unknown operator")
		//panic("Error")
	}
	return result
}
