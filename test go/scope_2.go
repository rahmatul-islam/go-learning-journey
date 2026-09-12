package main

import "fmt"

var x := 100

func main() {
	fmt.Println(x)

	x := 200
	p := x / 0
	fmt.Println(p)
	fmt.Println(x)
	var k int = ratul()
	fmt.Println(k)
}
func ratul() int {
	fmt.Println("Ratul")
	p := 500
	return p * p
}
