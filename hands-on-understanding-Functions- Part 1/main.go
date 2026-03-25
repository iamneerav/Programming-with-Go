/*
In this code, we will learn about diferent functions with examples

1. No parameters and No return type

2. With 1 parameters, No return type

3. With 1 parameter and 1 return type

4. With 2 parameters and 2 return type
*/

package main

import (
	"fmt"
)

func main() {

	greet()
	greetWithName("Neerav")

	greetWithNameAndReturn("Neerav")
	x := greetWithNameAndReturn("Neerav")
	fmt.Println(x)

	name, age := dogdetails("Buddy", 3)
	fmt.Println(name, age)

	sum(1, 2, 3, 4, 5)

	z := sum2(1, 2, 3, 4, 5)
	fmt.Println(z)

	Greeting("Hello", "Alice", "Bob")
	Greeting("nobody")
}

// 1. No parameters and No return type

func greet() {
	fmt.Println("Hello, World!")
}

// 2. With 1 parameters, No return type

func greetWithName(name string) {
	fmt.Printf("Hello, %v\n", name)
}

// 3. With 1 parameter and 1 return type

func greetWithNameAndReturn(name string) string {
	return fmt.Sprintln("Hello ", name)
}

// 4. With 2 parameters and 2 return type

func dogdetails(name string, age int) (string, int) {
	age += 1
	name = fmt.Sprintln("Dog's name is ", name)
	return name, age
}

/* Variadic Parameters

Variadic parameters allow a function to accept variable or any number of arguments.

1. A variadic parameter (...T) is treated like a slice ([]T) inside the function.
2. You can pass zero, one, or many strings here.

Example


func Greeting(prefix string, who ...string)

Use case 1: If we call > Greeting("nobody")

  So, Inside the function:

  	who becomes nil (meaning: “no slice, no values”)

	As we are calling "nobody" as prefix, so prefix will be "nobody" and who will be nil.


Use case 2: If we call > Greeting("Hello", "Alice", "Bob")

	So, inside the function

	"Hello" will be assigned to prefix, and who will be a slice of strings with the values ["Alice", "Bob"].

	As we are calling "Hello" as prefix, so prefix will be "Hello" and who will be ["Alice", "Bob"].

	Also, if we observe that the type of who is []string, which is a slice of strings. So,
	we can iterate over it using a loop or access its elements using an index.

https://go.dev/ref/spec#Passing_arguments_to_..._parameters

*/

// Example 1

func sum(nums ...int) {

	fmt.Println(nums)

	fmt.Printf("Variadic parameter type: %T\n", nums)
}

// Example 2

func sum2(nums ...int) int {
	i := 0
	for _, x := range nums {
		i += x
	}
	return i
}

// Example 3

func Greeting(prefix string, who ...string) {
	fmt.Printf("Prefix: %v\n", prefix)
	fmt.Printf("Who: %v\n", who)
	fmt.Printf("Type of who: %T\n", who)
}

/*
Unfurling a Slice

 Unfurling a slice is the process of passing a slice as individual arguments to a variadic function.

 func main() {

 a:= []int{1, 2, 3, 4, 5}

 x := sum(a...)  // We use the ... operator to unfurl the slice a and pass its elements as individual arguments to the sum function.

 fmt.Println(x)

 }


 func sum(num ...int) int {
	i := 0
	for _, x := range num {
		i += x
	}
	return i
}
*/

/*

Defer Method

The defer statement in Go is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.

// Example

func main() {

defer foo()
bar()

}

func foo() {

	fmt.Println("This is foo function")

}


func bar() {

	fmt.Println("This is bar function")

}


// Expected Output:

bar
foo

As, the defer statement defers the execution of the foo function until the surrounding function (main) returns. So, when we run the main function, it will first execute bar() and then, when main is about to return, it will execute foo.

*/

/*

Methods

In Go, a method is a function that is associated with a specific type. Methods allow you to define behavior for types, enabling you to perform operations on instances of those types. A method is defined with a receiver, which is the type it is associated with.

Example:

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

In this example, we have defined a struct type called Person with two fields: Name and Age. We have also defined a method called Greet that is associated with the Person type. The receiver of the method is (p Person), which means that the Greet method can be called on instances of the Person type.

To use the method, you can create an instance of the Person type and call the Greet method on it:

p1 := Person{
	Name: "Alice",
	Age:  30,
}

p2: = Person{
	Name: "Bob",
	Age:  25,
}

func main() {
	p1.Greet()
	p2.Greet()
}

This will output:

Hello, my name is Alice and I am 30 years old.
Hello, my name is Bob and I am 25 years old.

In this example, we created two instances of the Person type, p1 and p2, and called the Greet method on each of them. The Greet method uses the fields of the Person struct to print a greeting message.

*/
