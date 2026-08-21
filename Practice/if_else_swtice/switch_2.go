package main

import (
	"fmt"
	
)

func main(){

	var grate string

	fmt.Print("Enter your grate A to F:")
	fmt.Scan(&grate)

	switch grate{
	case "A","a":
		fmt.Print("Excellnt")
	case "B","b":
		fmt.Print("good")
	case "C","c":
		fmt.Print("Avarage")
	case "D","d":
		fmt.Print("Poor")
	case "F","f":
		fmt.Print("Fail")
	default:
		fmt.Print("you prace wrong letter")
	}
}