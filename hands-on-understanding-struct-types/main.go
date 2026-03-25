package main

import "fmt"


type person struct {
	Name string
	Age int
	Nationality string
}

type agentname struct {
	person 
	LicenseToKill bool
	Name string
}


func main() {

	// Normal Struct

	fmt.Println("Printing Normal Struct")

	x := person {
		Name: "Neerav",
		Age: 30,
		Nationality: "Indian",
	}

	fmt.Println(x)

	fmt.Println(x.Name)

	fmt.Println("---------")	

// Embedded Struct

	fmt.Println("Printing Embedded Struct")

	y := agentname {
		person: person {
			Name: "James Bond",
			Age: 40,
			Nationality: "British",
		},
		Name: "Neerav",
		LicenseToKill: true,
	}

	fmt.Println(y)
	
	// If we need to get Name field from agentName struct, we can do like this:
	fmt.Println(y.Name) // Accessing the Name field from agentname struct

	// If we need to get Name field from person struct, we can do like this:
	fmt.Println(y.person.Name) // Accessing the Name field from person struct

	// But if the we need to acces age, etc... we can do like this:
	fmt.Println(y.Age) // Accessing the Age field from person struct directly


// Anonymous Struct

/* 

1. We just take out the struct we want and have the details explicitly in the variable declaration.
2. Useful when we want to create a struct for one time use and don't want to define a new type for it.

*/

fmt.Println("Printing Anonymous Struct")

p3 := struct { 
	Name string
	Age int
	Nationality string
}{
	Name: "Neerav",
	Age: 30,
	Nationality: "Indian",
}

fmt.Printf("Type: %T \t Value: %v\n", p3, p3)




// Composition(Related to) Struct

fmt.Println("Printing Composition Struct")





}