/*
In this Functions exercise, we will learn about Callback
*/

/*

In a callback, we pass one function as an argument to another function.
The receiving function can call it later to perform dynamic behavior (like add or subtract).

*/

package main

import "fmt"

func main() {

	fmt.Println("-----Printing addition------")
	x := addnumber(5, 3)
	fmt.Println(x)

	fmt.Println("-----Printing subtraction------")
	y := subnumber(5, 3)
	fmt.Println(y)

	fmt.Println("-----Using calculator function with 'addnumber' as callback------")

	z := calculator(5, 3, addnumber)
	fmt.Println(z)

	fmt.Println("-----Using calculator function with 'subnumber' as callback------")

	a := calculator(5, 3, subnumber)
	fmt.Println(a)

}

func calculator(a int, b int, f func(a int, b int) int) int { // We are using a function as a parameter in calculator function. So, we have to specify the type of that function as well. In this case, it's func(a int, b int) int which means it takes two integers as input and returns an integer.
	return f(a, b)
}

func addnumber(a int, b int) int {
	return a + b
}

func subnumber(a int, b int) int {
	return a - b
}
