package main

import (
	"fmt"
)

func main() {

	// Calling the function 'foo' and 'bar'

	fmt.Println(foo(5))

	fmt.Println("-------------------")

	// Calling the function 'bar' and storing the returned values in variables

	x, s := bar(5, "Hello")
	fmt.Println(x, s)
}

func foo(x int) int {
	fmt.Printf("Return int\n")
	return x * 2
}

func bar(x int, a string) (int, string) {
	fmt.Printf("Return int and string\n")
	return x * 3, a

}
