package main

import "fmt"

type person struct {
	firstName  string
	lastName map[string]string
	favouriteIceCreamFlavours []string
}

func main() {


p1 := person {
	firstName: "Neerav",
	lastName: map[string]string{"lastName": "Bhatia"},
	favouriteIceCreamFlavours: []string{"Vanilla", "Chocolate", "Strawberry"},
}

p2 := person {
	firstName: "John",
	lastName: map[string]string{"lastName": "Smith"},
	favouriteIceCreamFlavours: []string{"Vanilla", "Nutella", "Pistachio"},
}

fmt.Println("Printing person struct")

fmt.Println(p1)

fmt.Println(p2)

fmt.Println("Printing flavours")

fmt.Println("Printing last names: Method 1")

for _, y := range p1.lastName {
	fmt.Println("My name is", p1.firstName, y)
	
}

for _, y := range p2.lastName {
	fmt.Println("My name is", p2.firstName, y)
	
}

}