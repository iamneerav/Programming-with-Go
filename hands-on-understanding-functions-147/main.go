/*
In this Functions exercise, we will learn about Returning a function
*/

/*

Returning a function means a function returns another function as a value.

In Go, that returned function can be stored in a variable and called later.
*/

package main

import "fmt"

func main() {

	x := foo()     // We call function foo and assign its return value to variable x
	fmt.Println(x) // We print the value of x which is '42'

	y := bar()     // We call function bar and assign its return value to variable y which is a function
	fmt.Println(y) // If we run like this, it will print the memory address of the function returned by bar. Similar to how we print a variable X.

	fmt.Println(y()) // But if we call the function y, it will execute the anonymous function and return 43

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

func foo() int { // We create a function foo that returns an int
	return 42
}

func bar() func() int { // We create a function bar that returns another function i.e. func() which further returns an int
	return func() int { // Returning an anonymous function that returns an int
		return 43 // Anonymous function that returns 43
	}
}
