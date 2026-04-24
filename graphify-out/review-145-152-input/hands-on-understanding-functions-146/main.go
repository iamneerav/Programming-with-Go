/*
In this Functions exercise, we will learn about func expression
*/

/*
In func expression, we can create a function and assign it to variable.
*/

package main

import "fmt"

func main() {

	foo() // We call normal function i.e. foo is func name followed by parentheses

	x := func() { // We then create function and assign it to variable x.
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) { // We then create function and assign it to variable y
		fmt.Println("This is an anonymous func showing my name", s)
	}

	x()         // x is a variable that holds the function
	y("Neerav") // y is a variable that holds the function and we pass argument "Neerav" to it
}

func foo() { // General func
	fmt.Println("Foo ran")
}

// a named function with an identifier
// func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func(p parameter(s)) (r return(s)) { code }
