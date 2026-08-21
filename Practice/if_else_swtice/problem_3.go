 package main

 import "fmt"

 func main(){
var mark int

fmt.Println("Enter your mark:")
fmt.Scan(&mark)

if mark>=40 {
	fmt.Println("pass")
	
}else{
	fmt.Println("fail")
}


 }