package main

import "fmt"

type person struct {
	firstName  string
	lastName string
	favouriteIceCreamFlavours []string
}

func main() {


p1 := person {
	firstName: "Neerav",
	lastName: "Bhatia",
	favouriteIceCreamFlavours: []string{"Vanilla", "Chocolate", "Strawberry"},
}

p2 := person {
	firstName: "John",
	lastName: "Smith",
	favouriteIceCreamFlavours: []string{"Vanilla", "Nutella", "Pistachio"},
}

fmt.Println("Printing person struct")

fmt.Println(p1)

fmt.Println(p2)

fmt.Println("Printing flavours")

for _, x := range p1.favouriteIceCreamFlavours {
	fmt.Println(p1.firstName, "favourite is", x)
}

for _, x := range p2.favouriteIceCreamFlavours {
	fmt.Println(p2.firstName, "favourite is", x)
}

}