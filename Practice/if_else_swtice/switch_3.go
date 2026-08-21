package main

import "fmt"

func main(){

	var a,b int
	fmt.Print("Enter first Number:")
	fmt.Scan(&a)
	fmt.Print("Enter second Number:")
	fmt.Scan(&b)

	var operator string
	fmt.Print("Enter the operator:")
	fmt.Scan(&operator)

	switch operator{
	case "+":
		fmt.Print("Sum is :",a+b)
	case "-":
		fmt.Print("sub is :",a-b)
	case "*":
		fmt.Print("multiplication is :",a*b)
	case "/":
		fmt.Print("division is :",a/b)
	default:
		fmt.Print("you type wrong sing")
	}
	




}