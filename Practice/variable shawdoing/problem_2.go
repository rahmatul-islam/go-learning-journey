
/*
package main

import "fmt"

func calculate(x int) {
	x := 100 // Error: no new variables on left side of :=
	fmt.Println(x)
}

func main() {
	calculate(50)
}
*/

package main

import "fmt"

func calculate(x int) {
	fmt.Println("Parameter x:", x)

	if true {
		x := 100
		fmt.Println("Shadowed x:", x)
	}

	fmt.Println("Original x:", x)
}

func main() {
	calculate(50)
}