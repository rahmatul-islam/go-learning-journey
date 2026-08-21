package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Print("Enter 1st number:")
	fmt.Scan(&num1)
	fmt.Print("Enter 2nd number:")
	fmt.Scan(&num2)
	fmt.Print("Enter 3rd number:")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {

		fmt.Print("1s number is big")
	} else if num1 < num2 && num2 > num3 {

		fmt.Print("2n number is big")

	} else {
		fmt.Print("3rd number is big")
	}
}
