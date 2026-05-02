package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(namedreturn(a))

}

func namedreturn(x []int) (sum int) { // sum int is a named return variable, which means that it is declared in the function signature and can be used in the function body without being explicitly declared. The value of sum will be returned when the function returns, even if there is no explicit return statement.

	sum = 0
	for _, i := range x {

		sum += i

	}
	return sum
}

// Another way of writing the above function is:

/*

func namedreturn(x []int) int {

	sum := 0
	for _, i := range x {

		sum += i

	}
	return sum
}
*/
