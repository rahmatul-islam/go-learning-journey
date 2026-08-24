package main
import"fmt"

func getNumber(num int){
	if num>0 {
		fmt.Print("positive")
	} else if num<0 {
		fmt.Print("Negative")
	} else {
		fmt.Print("zero")
	}
}

func main(){

	getNumber(10)
}