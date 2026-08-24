package main

import "fmt"

func getinput() (float64, float64, string) {
	var a, b float64
	var operator string

	fmt.Print("Enter first Number:")
	fmt.Scan(&a)
	fmt.Print("Enter second Number:")
	fmt.Scan(&b)
	fmt.Print("Enter opretion:")
	fmt.Scan(&operator)

	return a, b, operator

}

func calculator(a, b float64, operator string) float64 {

	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func main() {
	a, b, operator := getinput()

	fmt.Println(calculator(a, b, operator))
}
