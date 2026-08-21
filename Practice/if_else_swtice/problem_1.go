package main

import"fmt"
func main(){
	var num int

	fmt.Print("Enter intiger number:")
	fmt.Scan(&num)

	if num%2==0 {
		fmt.Println("Number is even")
		
	}else {
		fmt.Println("Number is odd")
	}

}