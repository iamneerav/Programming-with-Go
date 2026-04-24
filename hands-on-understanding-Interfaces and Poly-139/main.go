/*
In this Functions exercise, we will learn about the following topics:

1.Interfaces and Polymorphism

*/

/*

Revision of Function, Types and Methods

*/

package main

import (
	"fmt"
)

func main() {

	// Calling the function 'doggie'

	doggie("Goovie")

	fmt.Println("-------------------")

	// Creating different instances of Dog Struct

	d1 := Dog{
		Name:  "Charlie",
		Age:   3,
		Breed: "Golden Retriever",
	}

	d2 := Dog{
		Name:  "Poochie",
		Age:   5,
		Breed: "Poodle",
	}

	// Calling the methods for different instances of Dog Struct

	d1.details()
	d1.bark()

	fmt.Println("-------------------")

	d2.details()
	d2.bark()

	// Calling instances of Owner Struct

	fmt.Println("-------------------")

	o1 := Owner{
		Dog: Dog{
			Name: "Charlie",
		},
		Age: 30,
	}

	// Calling the methods for instances of Owner Struct

	o1.details()

	// Using Polymorphism and Interfaces

	fmt.Println("-------------------")

	pethumaninformation(d1) // Will call the `details` method of Dog struct
	pethumaninformation(o1) // Will call the `details` method of Owner struct
}

func doggie(x string) {
	fmt.Println("My name is ", x)
}

// Using Methods (Can be used with different types)

// Implementing Methods

/// Creating struct

type Dog struct {
	Name  string
	Age   int
	Breed string
}

// Creating different methods for Type Dog Struct

func (d Dog) bark() {
	fmt.Println("I do Woof Woof")
}

func (d Dog) details() {
	fmt.Println("My name is ", d.Name, ". I am ", d.Age, " years old and I am a ", d.Breed)

}

/*

Polymorphism and Interfaces

// Definitions

Interface: Defines a set of method signatures that a type must implement to satisfy the interface. It allows you to define behavior without specifying the underlying implementation.

Polymorphism: The ability of different types to be treated as instances of the same interface. It allows you to write code that can work with different types that implement the same interface.

*/

// Now we added another struct

type Owner struct {
	Dog
	Age int
}

func (o Owner) details() {
	fmt.Println(" I am the owner of", o.Name, "and I live at Rock street, New York")

}

// Lets define an interface

// An interface is a collection of method signatures. This has details method signature. So, any 'type' that has this method signature will satisfy this interface and can be treated as an instance of this interface. In this case, both Dog and Owner struct have details method signature. So, both of them will satisfy this interface and can be treated as an instance of this interface.

type Pet interface {
	details()
}

// Lets create a function that takes an instance of Pet interface as an argument. This function can accept any type that satisfies the Pet interface, which means it can accept both Dog and Owner struct instances.

func pethumaninformation(p Pet) {
	p.details()
}
