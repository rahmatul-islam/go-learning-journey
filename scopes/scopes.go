package main
import "fmt"
var p int = 100
func main(){
	fmt.Println(p)
	p:=200
	fmt.Println(p)
	fmt.Println("Ratul The Danger Boy")
	k:=ratul()
	y:=p+k
	fmt.Println("Ratul's Age+Money in pocket is:=",y)
}
func ratul() int {
	age:=21
	return age
}