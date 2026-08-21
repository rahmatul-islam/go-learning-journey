

package main 

import"fmt"

func main(){

	var Number int;

	fmt.Print("Enter a number 1 To 7:")
	fmt.Scan(&Number)

	switch Number{
	case 1:
		fmt.Print("Monday")
	case 2:
		fmt.Print("Tuseday")
	case 3:
		fmt.Print("Wednesday")
	case 4:
		fmt.Print("Thursday")
	case 5:
		fmt.Print("Friday")
	case 6:
		fmt.Print("Saturday")
	case 7:
		fmt.Print("Sunday")
	default:
		fmt.Print("You enter Worng number")	
	}

}



