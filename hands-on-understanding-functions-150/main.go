/*
In this Functions exercise, we will learn about Function Fundamentals

*/

/*

Basic function format in Go:

func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

- Receiver is used in methods, we will learn about it later. For now, we can ignore it.
- Identifier is the name of the function.
- Parameter(s) are the input values that the function can take. They are optional, a function can have no parameters as well.
- Return(s) are the output values that the function can return. They are also optional, a function can have no return value as well.

*/

package main

import "fmt"

func main() {
	foo()
}

func foo() {
	fmt.Println("Foo ran")
}
