package main

import "fmt"

func main() {

	var age int

	fmt.Print("Enter your age:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Print("you can vote")

	} else {
		fmt.Print("you can not vote")
	}

}
