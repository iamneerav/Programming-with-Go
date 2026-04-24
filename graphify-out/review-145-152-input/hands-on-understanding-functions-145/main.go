/*
In this Functions exercise, we will learn about Anonymous Function
*/

/*

We have function in the format:

func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// This is an anonymous function: There is NO IDENTIFIER

func(p parameter(s)) (r return(s)) { code }

*/

package main

import "fmt"

func main() {

	foo()

	// This is an anonymous func

	// Example 1:

	func() { // We provide func(){}() --> Why () at the end? Because this is where we pass arguments in func (main), whenever it needs input.
		fmt.Println("Anonymous func ran")
	}()

	// Example 2:

	func(s string) { // We create anony. func with one parameter as an input.
		fmt.Println("This is an anonymous func showing my name", s)
	}("Neerav") // So, now if we carefully see this, it needs one input as an argument.

}

// Example below of general function

func foo() { // This is just a general 'foo' func to demonstrate difference between regular and anonymous func.
	fmt.Println("Foo ran")
}
