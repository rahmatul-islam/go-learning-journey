package main

import "fmt"

func main (){

	var num int

	fmt.Print("Enter you Number:")
	fmt.Scan(&num)

	if num%5==0 {
		fmt.Print("Yes")
		
	}else{
		fmt.Print("No")
	}


}