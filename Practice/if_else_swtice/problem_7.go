package main

import "fmt"

func main(){
 var amount int
 fmt.Print("Enter your amount")
 fmt.Scan(&amount)

 if amount>=1000 {
	fmt.Print("10% discount")
	
 }else{
	fmt.Print("No discount")
 }

}