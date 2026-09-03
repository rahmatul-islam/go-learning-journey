
package main

import "fmt"

func main() {
	x := 10

	{
		y := 20
		fmt.Println(x)
		fmt.Println(y)
	}

	fmt.Println(x)
	fmt.Println(y)
}