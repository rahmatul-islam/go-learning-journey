package main

import "fmt"

func Input() (int, int) {
	var firstNumber, secondNumber int

	fmt.Print("Enter first Number:")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second Number:")
	fmt.Scan(&secondNumber)
	return firstNumber, secondNumber
}

func bigNumber(firstNumber, secondNumber int) int {
	if firstNumber > secondNumber {
		return firstNumber
	}
	return secondNumber
}

func main() {
	firstNumber, secondNumber := Input()
	fmt.Println("Bigger number:", bigNumber(firstNumber, secondNumber))
}
