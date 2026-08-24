package main

import "fmt"
 func checkPassword(password string) bool{

	if password == "golang123" {
		
		return true
	}else{
		return false
	}
}
func getinput()string{

	var password string
	fmt.Print("Enter your password:")
	fmt.Scan(&password)
	return password
}
func main(){

	password :=getinput()
	if checkPassword(password)==true {
		fmt.Println("login successfull")
	}else{
		fmt.Print("Login faild")
	}
}