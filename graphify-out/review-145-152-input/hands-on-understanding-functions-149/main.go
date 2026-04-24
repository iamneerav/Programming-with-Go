/*
In this Functions exercise, we will learn about Closure
*/

/*
If we think about it, a closure is a function that "closes over" from its surrounding scope.


*/

package main

import "fmt"

func main() {
	f := incrementor()

	fmt.Println("Calling f which will print the memory address of the function returned by incrementor and it will show the type of the function")

	fmt.Printf("Value of f: %v\n Type: %T\n", f, f) // This will print the memory address of the function returned by incrementor. Similar to how we print a variable X.

	// The below will execute the anonymous function returned by incrementor and return 1. This is because the variable x in incrementor starts at 0 and then gets incremented by 1 before being returned.

	fmt.Println("Creating a closure with variable f and this closure will have access to variable x defined in incrementor function")

	fmt.Printf("Value of f: %v\n Type: %T\n", f(), f())

	fmt.Println(f())

	// Similarly, the next calls to f() will return 2, 3, 4, 5, and 6 as it keeps incrementing x by 1 each time.

	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6

	// Now, since f was assigned before and it was stored in memory, but now we will see an example where we define a new variable g and assign it the value of incrementor() which will create a new closure with its own variable x starting at 0.

	fmt.Println("Creating a new closure with variable g")

	g := incrementor()
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	fmt.Println(g()) // 4
	fmt.Println(g()) // 5
	fmt.Println(g()) // 6
}

func incrementor() func() int { // Creating a function 'incrementor' that returns another function which is of type func() int
	x := 0              // This will return x as 0 when we call incrementor for the first time.
	return func() int { // This will return an anonymous function that returns an int. This is the closure part where the inner function has access to the variable x defined in the outer function incrementor.
		x++
		return x
	}
}
