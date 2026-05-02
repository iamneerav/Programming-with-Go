package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello! My name is %v and I am %v years old.\n", p.first, p.age)
}

func main() {

	p1 := person{
		first: "Neerav",
		age:   29,
	}

	p1.speak()
}
