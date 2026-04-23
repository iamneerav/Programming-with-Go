/*
In this Functions exercise, we will learn about Recursion
*/

/*

Recursion in programming refers to the technique of solving a problem by breaking it down into smaller, more manageable subproblems. A recursive function is a function that calls itself in order to solve a problem.

The function continues to call itself with modified arguments until it reaches a **base** case, which is a condition that stops the recursion.
*/

/*

We will take a simple example of calculating the factorial of a number using recursion.

For Example:

5! = 5 * 4 * 3 * 2 * 1 = 120

which can be expressed as:

factorial(n) = n * n-1 * n-2 * ... * 1

or in terms of recursion:

factorial(n) = n * factorial(n-1)

Example of factorial function in Go:

factorial(5) = 5 * factorial(4)
factorial(4) = 4 * factorial(3)
factorial(3) = 3 * factorial(2)
factorial(2) = 2 * factorial(1)
factorial(1) = 1 (base case)	-> This is where the recursion stops.

So, the function will keep calling itself until it reaches the base case where n is 1, at which point it will return 1 and then the previous calls will start returning their results until we get the final result of 120 for factorial(5).

*/

package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(a int) int {
	if a == 1 {
		return 1
	}
	a = a * factorial(a-1)
	return a
}
