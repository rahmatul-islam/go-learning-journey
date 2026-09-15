package main
 
import"fmt"
// global variable
 var x=10
//cerate function
func slow(){

    x:=20
	fmt.Println(x)
}
func main(){
	
// call function
	slow()
	fmt.Println(x)
}